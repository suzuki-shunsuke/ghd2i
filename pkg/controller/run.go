package controller

import (
	"bytes"
	"context"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/shurcooL/githubv4"
	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/ghd2i/pkg/github"
	"github.com/suzuki-shunsuke/logrus-error/logerr"
)

type Param struct {
	ConfigFilePath string
	DataFilePath   string
	Close          string
	Lock           string
	RepoOwner      string
	RepoName       string
	Query          string
	Args           []string
	Labels         []string
	Assignees      []string
	DryRun         bool
}

type Config struct {
	Title           string
	IssueTemplate   string `yaml:"issue_template"`
	CommentTemplate string `yaml:"comment_template"`
}

type GitHub interface {
	GetDiscussion(ctx context.Context, owner, name string, number int) (*github.Discussion, error)
	CreateIssue(ctx context.Context, owner, repo string, req *github.IssueRequest) (int, string, error)
	CreateIssueComment(ctx context.Context, owner, name string, number int, req *github.IssueComment) (string, error)
	MinimizeComment(ctx context.Context, nodeID string, minimizedReason githubv4.ReportedContentClassifiers) error
	LockIssue(ctx context.Context, owner, name string, number int, lockReason string) error
	CloseIssue(ctx context.Context, owner, name string, number int) error
	SearchDiscussions(ctx context.Context, query string) ([]string, error)
}

type ParamDiscussion struct {
	Owner  string
	Name   string
	Number int
}

//go:embed issue_body.tpl
var issueBodyTplByte []byte

//go:embed issue_comment.tpl
var issueCommentBodyTplByte []byte

func (c *Controller) Run(ctx context.Context, logE *logrus.Entry, param *Param) error { //nolint:cyclop
	if param.Query == "" && param.DataFilePath == "" && len(param.Args) == 0 {
		logE.Warn("no discussion is specified. One of -query, -data, or discussion URLs are required")
		return nil
	}
	cfg := &Config{}
	if err := findAndReadConfig(c.fs, cfg, param.ConfigFilePath); err != nil {
		return fmt.Errorf("find and read a configuration file: %w", err)
	}
	if err := c.parseTemplates(cfg); err != nil {
		return err
	}
	failed := false
	if param.DataFilePath != "" {
		// read data file
		discussions := &Discussions{}
		if err := c.readData(param.DataFilePath, discussions); err != nil {
			return fmt.Errorf("read a data file: %w", err)
		}
		for _, discussion := range discussions.Discussions {
			if err := c.run(ctx, logE, param, discussion); err != nil {
				logerr.WithError(logE, err).Error("handle a discussion")
				failed = true
			}
		}
		if failed {
			return errors.New("failed to handle some discussions")
		}
	}
	args := param.Args
	if param.Query != "" {
		// search discussions by GitHub GraphQL API
		urls, err := c.gh.SearchDiscussions(ctx, "is:discussions "+param.Query)
		if err != nil {
			return fmt.Errorf("search discussions: %w", err)
		}
		args = make([]string, len(param.Args), len(param.Args)+len(urls))
		copy(args, param.Args)
		args = append(args, urls...)
	}
	for _, arg := range args {
		logE := logE.WithField("arg", arg)
		discussion, err := c.getDiscussion(ctx, arg)
		if err != nil {
			logerr.WithError(logE, err).Error("get a discussion")
			failed = true
		}
		if err := c.run(ctx, logE, param, discussion); err != nil {
			logerr.WithError(logE, err).Error("handle a discussion")
			failed = true
		}
	}
	if failed {
		return errors.New("failed to handle some discussions")
	}
	return nil
}

func (c *Controller) parseTemplates(cfg *Config) error {
	if cfg.IssueTemplate != "" {
		t, err := parseTemplate(cfg.IssueTemplate)
		if err != nil {
			return fmt.Errorf("parse an issue template in the configuration file: %w", err)
		}
		c.issueBody = t
	}
	if cfg.CommentTemplate != "" {
		cmt, err := parseTemplate(cfg.CommentTemplate)
		if err != nil {
			return fmt.Errorf("parse a comment template in the configuration file: %w", err)
		}
		c.issueCommentBody = cmt
	}
	if cfg.Title != "" {
		title, err := parseTemplate(cfg.Title)
		if err != nil {
			return fmt.Errorf("parse a title template in the configuration file: %w", err)
		}
		c.title = title
	}
	return nil
}

var discussionURLPattern = regexp.MustCompile(`https://github.com/([^/]+)/([^/]+)/discussions/(\d+)`)

func parseArg(arg string) (*ParamDiscussion, error) {
	// https://github.com/suzuki-shunsuke/test-github-action/discussions/55
	arr := discussionURLPattern.FindStringSubmatch(arg)
	if arr == nil {
		return nil, errors.New("arg must be a GitHub Discussion URL: " + arg)
	}
	if len(arr) != 4 { //nolint:mnd
		return nil, errors.New("arg must be a GitHub Discussion URL: " + arg)
	}
	number, err := strconv.Atoi(arr[3])
	if err != nil {
		return nil, fmt.Errorf("parse a discussion number as a number: %w", err)
	}
	return &ParamDiscussion{
		Owner:  arr[1],
		Name:   arr[2],
		Number: number,
	}, nil
}

func (c *Controller) run(ctx context.Context, logE *logrus.Entry, param *Param, discussion *Discussion) error { //nolint:funlen,cyclop,gocognit
	// Render issue and comments based on templates.
	buf := &bytes.Buffer{}
	if err := c.issueBody.Execute(buf, discussion); err != nil {
		return fmt.Errorf("render an issue body using a template engine: %w", err)
	}
	comments := make([]string, len(discussion.Comments))
	for i, comment := range discussion.Comments {
		buf := &bytes.Buffer{}
		if err := c.issueCommentBody.Execute(buf, comment); err != nil {
			return fmt.Errorf("render an issue comment body using a template engine: %w", err)
		}
		comments[i] = buf.String()
	}
	issueBody := buf.String()
	if param.DryRun {
		arr := append([]string{issueBody}, comments...)
		fmt.Fprintln(c.stdout, strings.Join(arr, "\n\n---\n\n"))
		return nil
	}
	title := discussion.Title
	if c.title != nil {
		buf := &bytes.Buffer{}
		if err := c.title.Execute(buf, discussion); err != nil {
			return fmt.Errorf("render an issue title using a template engine: %w", err)
		}
		title = buf.String()
	}
	repoOwner := discussion.Repo.Owner
	repoName := discussion.Repo.Name
	if param.RepoOwner != "" {
		repoOwner = param.RepoOwner
	}
	if param.RepoName != "" {
		repoName = param.RepoName
	}
	// Create an issue by GitHub GraphQL API.
	labels := make([]string, len(discussion.Labels), len(discussion.Labels)+len(param.Labels))
	copy(labels, discussion.Labels)
	labels = append(labels, param.Labels...)
	issueRequest := &github.IssueRequest{
		Title:  &title,
		Body:   &issueBody,
		Labels: &labels,
	}
	if param.Assignees != nil {
		issueRequest.Assignees = &param.Assignees
	}
	issueNum, issueURL, err := c.gh.CreateIssue(ctx, repoOwner, repoName, issueRequest)
	if err != nil {
		return fmt.Errorf("create an Issue: %w", err)
	}
	logE.WithField("issue_url", issueURL).Info("created an issue")
	// Create issue comments by GitHub GraphQL API.
	for i, comment := range discussion.Comments {
		commentID, err := c.gh.CreateIssueComment(ctx, repoOwner, repoName, issueNum, &github.IssueComment{
			Body: &comments[i],
		})
		if err != nil {
			return fmt.Errorf("create a comment: %w", err)
		}
		if comment.IsMinimized {
			reason, ok := github.GetMinimizedReason(comment.MinimizedReason)
			if !ok {
				logE.WithField("minimized_reason", comment.MinimizedReason).Warn("unknown minimized reason")
				reason = githubv4.ReportedContentClassifiersResolved
			}
			if err := c.gh.MinimizeComment(ctx, commentID, reason); err != nil {
				logerr.WithError(logE, err).WithField("comment_id", commentID).Warn("minimize a comment")
			}
		}
	}
	// Close and lock the issue if necessary.
	if param.Close == "always" || discussion.Closed && param.Close != "never" {
		if err := c.gh.CloseIssue(ctx, repoOwner, repoName, issueNum); err != nil {
			return fmt.Errorf("close an issue: %w", err)
		}
	}
	if param.Lock == "always" || discussion.Locked && param.Lock != "never" {
		if err := c.gh.LockIssue(ctx, repoOwner, repoName, issueNum, "resolved"); err != nil {
			return fmt.Errorf("lock an issue: %w", err)
		}
	}
	return nil
}

func (c *Controller) readData(dataFilePath string, discussions *Discussions) error {
	// read data file
	f, err := c.fs.Open(dataFilePath)
	if err != nil {
		return fmt.Errorf("open a data file: %w", err)
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(discussions); err != nil {
		return fmt.Errorf("read a data file as JSON: %w", err)
	}
	return nil
}
