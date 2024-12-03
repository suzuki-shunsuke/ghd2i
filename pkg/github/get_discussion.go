package github

import (
	"context"
	"fmt"

	"github.com/shurcooL/githubv4"
)

func (c *Client) GetDiscussion(ctx context.Context, owner, name string, number int) (*Discussion, error) {
	q := &Query{}
	variables := map[string]any{
		"repoOwner": githubv4.String(owner),
		"repoName":  githubv4.String(name),
		"number":    githubv4.Int(number), //nolint:gosec
	}
	if err := c.v4Client.Query(ctx, &q, variables); err != nil {
		return nil, fmt.Errorf("get a discussion by GitHub GraphQL API: %w", err)
	}
	return q.Repository.Discussion, nil
}
