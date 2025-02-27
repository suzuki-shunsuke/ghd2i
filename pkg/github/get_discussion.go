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

	if q.Repository.Discussion.Comments.PageInfo.HasNextPage {
		for range 10 {
			comments, err := c.SearchComments(ctx, owner, name, number, q.Repository.Discussion.Comments.PageInfo.EndCursor)
			if err != nil {
				return nil, fmt.Errorf("search discussion comments by GitHub GraphQL API: %w", err)
			}
			q.Repository.Discussion.Comments.Nodes = append(q.Repository.Discussion.Comments.Nodes, comments...)
		}
	}
	if q.Repository.Discussion.Reactions.PageInfo.HasNextPage {
		for range 10 {
			reactions, err := c.SearchDiscussionReactions(ctx, owner, name, number, q.Repository.Discussion.Comments.PageInfo.EndCursor)
			if err != nil {
				return nil, fmt.Errorf("search discussion reactions by GitHub GraphQL API: %w", err)
			}
			q.Repository.Discussion.Reactions.Nodes = append(q.Repository.Discussion.Reactions.Nodes, reactions...)
		}
	}
	return q.Repository.Discussion, nil
}

func (c *Client) SearchComments(ctx context.Context, owner, name string, number int, cursor string) ([]*Comment, error) { //nolint:dupl
	var comments []*Comment
	variables := map[string]any{
		"repoOwner": githubv4.String(owner),
		"repoName":  githubv4.String(name),
		"number":    githubv4.Int(number), //nolint:gosec
		"cursor":    cursor,
	}
	for range 100 {
		q := &SearchCommentQuery{}
		if err := c.v4Client.Query(ctx, q, variables); err != nil {
			return nil, fmt.Errorf("search discussion comments by GitHub GraphQL API: %w", err)
		}
		comments = append(comments, q.Repository.Discussion.Comments.Nodes...)
		if !q.Repository.Discussion.Comments.PageInfo.HasNextPage {
			return comments, nil
		}
		variables["cursor"] = q.Repository.Discussion.Comments.PageInfo.HasNextPage
	}
	return nil, nil
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

func (c *Client) SearchDiscussionReactions(ctx context.Context, owner, name string, number int, cursor string) ([]*Reaction, error) { //nolint:dupl
	var reactions []*Reaction
	variables := map[string]any{
		"repoOwner": githubv4.String(owner),
		"repoName":  githubv4.String(name),
		"number":    githubv4.Int(number), //nolint:gosec
		"cursor":    cursor,
	}
	for range 10 {
		q := &SearchDiscussionReactionsQuery{}
		if err := c.v4Client.Query(ctx, q, variables); err != nil {
			return nil, fmt.Errorf("search discussion reactions by GitHub GraphQL API: %w", err)
		}
		reactions = append(reactions, q.Repository.Discussion.Reactions.Nodes...)
		if !q.Repository.Discussion.Reactions.PageInfo.HasNextPage {
			return reactions, nil
		}
		variables["cursor"] = q.Repository.Discussion.Reactions.PageInfo.HasNextPage
	}
	return nil, nil
}
