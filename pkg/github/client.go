package github

import (
	"context"

	"github.com/google/go-github/v69/github"
	"github.com/shurcooL/githubv4"
	"golang.org/x/oauth2"
)

type Client struct {
	v3Client *github.Client
	v4Client *githubv4.Client
}

func New(ctx context.Context, token string) *Client {
	httpClient := oauth2.NewClient(ctx, oauth2.StaticTokenSource(
		&oauth2.Token{AccessToken: token},
	))

	return &Client{
		v3Client: github.NewClient(httpClient),
		v4Client: githubv4.NewClient(httpClient),
	}
}
