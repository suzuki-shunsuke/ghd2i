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
	var urls []string
	variables := map[string]any{
		"query":  githubv4.String(query),
		"cursor": (*githubv4.String)(nil),
	}
	for range 10 {
		q := &SearchQuery{}
		if err := c.v4Client.Query(ctx, q, variables); err != nil {
			return nil, fmt.Errorf("search discussions by GitHub GraphQL API: %w", err)
		}
		for _, node := range q.Search.Nodes {
			urls = append(urls, node.Discussion.URL)
		}
		if !q.Search.PageInfo.HasNextPage {
			return urls, nil
		}
		variables["cursor"] = q.Search.PageInfo.EndCursor
	}
	return urls, nil
}
