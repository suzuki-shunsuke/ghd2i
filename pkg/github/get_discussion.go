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
	if err := c.v4Client.Query(ctx, q, variables); err != nil {
		return nil, fmt.Errorf("get a discussion by GitHub GraphQL API: %w", err)
	}
	return q.Repository.Discussion, nil
}

func (c *Client) SearchDiscussions(ctx context.Context, query string) ([]string, error) {
	q := &SearchQuery{}
	variables := map[string]any{
		"query": githubv4.String(query),
	}
	if err := c.v4Client.Query(ctx, q, variables); err != nil {
		return nil, fmt.Errorf("search discussions by GitHub GraphQL API: %w", err)
	}
	urls := make([]string, len(q.Search.Nodes))
	for i, node := range q.Search.Nodes {
		urls[i] = node.Discussion.URL
	}
	return urls, nil
}
