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
		return 0, "", err
	}
	return issue.GetNumber(), issue.GetHTMLURL(), nil
}

func (c *Client) CreateIssueComment(ctx context.Context, owner, name string, number int, req *github.IssueComment) (string, error) {
	// https://pkg.go.dev/github.com/google/go-github/v67/github#IssuesService.CreateComment
	comment, _, err := c.v3Client.Issues.CreateComment(ctx, owner, name, number, req)
	if err != nil {
		return "", err
	}
	return comment.GetNodeID(), nil
}

func (c *Client) MinimizeComment(ctx context.Context, nodeID string) error {
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
		Classifier: githubv4.ReportedContentClassifiersOutdated,
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
		return err
	}
	return nil
}

func (c *Client) LockIssue(ctx context.Context, owner, name string, number int, lockReason string) error {
	// https://pkg.go.dev/github.com/google/go-github/v67/github#IssuesService.Lock
	_, err := c.v3Client.Issues.Lock(ctx, owner, name, number, &github.LockIssueOptions{
		LockReason: lockReason,
	})
	if err != nil {
		return err
	}
	return nil
}
