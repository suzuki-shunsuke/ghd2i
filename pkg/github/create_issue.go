package github

import (
	"context"
	"fmt"

	"github.com/google/go-github/v67/github"
	"github.com/shurcooL/githubv4"
)

type (
	IssueRequest = github.IssueRequest
	IssueComment = github.IssueComment
)

func String(s string) *string {
	return &s
}

func (c *Client) CreateIssue(ctx context.Context, owner, repo string, req *github.IssueRequest) (int, string, error) {
	// https://pkg.go.dev/github.com/google/go-github/v67/github#IssuesService.Create
	// https://docs.github.com/en/rest/issues/issues#create-an-issue
	issue, _, err := c.v3Client.Issues.Create(ctx, owner, repo, req)
	if err != nil {
		return 0, "", err //nolint:wrapcheck
	}
	return issue.GetNumber(), issue.GetHTMLURL(), nil
}

func (c *Client) CreateIssueComment(ctx context.Context, owner, name string, number int, req *github.IssueComment) (string, error) {
	// https://pkg.go.dev/github.com/google/go-github/v67/github#IssuesService.CreateComment
	comment, _, err := c.v3Client.Issues.CreateComment(ctx, owner, name, number, req)
	if err != nil {
		return "", err //nolint:wrapcheck
	}
	return comment.GetNodeID(), nil
}

func GetMinimizedReason(reason string) (githubv4.ReportedContentClassifiers, bool) {
	reasons := map[string]githubv4.ReportedContentClassifiers{
		"spam":      githubv4.ReportedContentClassifiersSpam,
		"abuse":     githubv4.ReportedContentClassifiersAbuse,
		"off-topic": githubv4.ReportedContentClassifiersOffTopic,
		"outdated":  githubv4.ReportedContentClassifiersOutdated,
		"duplicate": githubv4.ReportedContentClassifiersDuplicate,
		"resolved":  githubv4.ReportedContentClassifiersResolved,
	}
	a, ok := reasons[reason]
	return a, ok
}

func (c *Client) MinimizeComment(ctx context.Context, nodeID string, minimizedReason githubv4.ReportedContentClassifiers) error {
	// minimizeComment
	// https://docs.github.com/en/graphql/reference/mutations#minimizecomment
	var m struct {
		MinimizeComment struct {
			MinimizedComment struct {
				MinimizedReason   githubv4.String
				IsMinimized       githubv4.Boolean
				ViewerCanMinimize githubv4.Boolean
			}
		} `graphql:"minimizeComment(input:$input)"`
	}

	input := githubv4.MinimizeCommentInput{
		Classifier: minimizedReason,
		SubjectID:  nodeID,
	}
	if err := c.v4Client.Mutate(ctx, &m, input, nil); err != nil {
		return fmt.Errorf("minimize an issue comment: %w", err)
	}
	return nil
}

func (c *Client) CloseIssue(ctx context.Context, owner, name string, number int) error {
	// closeIssue
	// https://docs.github.com/en/graphql/reference/mutations#closeissue
	_, _, err := c.v3Client.Issues.Edit(ctx, owner, name, number, &github.IssueRequest{
		State: String("closed"),
	})
	if err != nil {
		return err //nolint:wrapcheck
	}
	return nil
}

func (c *Client) LockIssue(ctx context.Context, owner, name string, number int, lockReason string) error {
	// https://pkg.go.dev/github.com/google/go-github/v67/github#IssuesService.Lock
	_, err := c.v3Client.Issues.Lock(ctx, owner, name, number, &github.LockIssueOptions{
		LockReason: lockReason,
	})
	if err != nil {
		return err //nolint:wrapcheck
	}
	return nil
}

func (c *Client) CloseDiscussion(ctx context.Context, id string, reason githubv4.DiscussionCloseReason) error {
	var m struct {
		Discussion struct {
			ID string
		} `graphql:"closeDiscussion(input:$input)"`
	}

	input := githubv4.CloseDiscussionInput{
		DiscussionID: id,
		Reason:       &reason,
	}
	if err := c.v4Client.Mutate(ctx, &m, input, nil); err != nil {
		return fmt.Errorf("close a discussion: %w", err)
	}
	return nil
}

func (c *Client) LockDiscussion(ctx context.Context, id string) error {
	var m struct {
		Discussion struct {
			ID string
		} `graphql:"lockLockable(input:$input)"`
	}

	input := githubv4.LockLockableInput{
		LockableID: id,
	}
	if err := c.v4Client.Mutate(ctx, &m, input, nil); err != nil {
		return fmt.Errorf("lock a discussion: %w", err)
	}
	return nil
}
