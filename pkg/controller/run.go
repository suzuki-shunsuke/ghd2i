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

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/ghd2i/pkg/github"
)

type Param struct {
	ConfigFilePath string
	DataFilePath   string
	Args           []string
	DryRun         bool
}

type Config struct{}

type GitHub interface {
	GetDiscussion(ctx context.Context, owner, name string, number int) (*github.Discussion, error)
	// CreateIssue
	// CreateIssueComment
	// HideComment
	// CloseIssue
	// LockIssue
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

func (c *Controller) run(ctx context.Context, logE *logrus.Entry, param *Param, discussion *Discussion) error {
	// Render issue and comments based on templates.
	buf := &bytes.Buffer{}
	if err := c.issueBody.Execute(buf, discussion); err != nil {
		return err
	}
	comments := make([]string, len(discussion.Comments))
	for i, comment := range discussion.Comments {
		buf := &bytes.Buffer{}
		if err := c.issueCommentBody.Execute(buf, comment); err != nil {
			return err
		}
		comments[i] = buf.String()
	}
	if param.DryRun {
		arr := append([]string{buf.String()}, comments...)
		fmt.Fprintln(c.stdout, strings.Join(arr, "\n\n---\n\n"))
		return nil
	}
	// Create an issue by GitHub GraphQL API.
	// Create issue comments by GitHub GraphQL API.
	// Hide comments if necessary.
	// Close and lock the issue if necessary.
	return nil
}

func (c *Controller) readData(dataFilePath string, discussions *Discussions) error {
	// read data file
	f, err := c.fs.Open(dataFilePath)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := json.NewDecoder(f).Decode(discussions); err != nil {
		return err
	}
	return nil
}

func (c *Controller) Run(ctx context.Context, logE *logrus.Entry, param *Param) error {
	if param.DataFilePath != "" {
		// read data file
		discussions := &Discussions{}
		if err := c.readData(param.DataFilePath, discussions); err != nil {
			return err
		}
		for _, discussion := range discussions.Discussions {
			if err := c.run(ctx, logE, param, discussion); err != nil {
				return err
			}
		}
		return nil
	}
	for _, arg := range param.Args {
		discussion, err := c.getDiscussion(ctx, arg)
		if err != nil {
			return err
		}
		if err := c.run(ctx, logE, param, discussion); err != nil {
			return err
		}
	}
	return nil
}
