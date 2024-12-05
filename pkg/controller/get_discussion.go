package controller

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/ghd2i/pkg/github"
)

func (c *Controller) GetDiscussion(ctx context.Context, _ *logrus.Entry, param *Param) error {
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
	discussions := make([]*Discussion, len(args))
	for i, arg := range args {
		d, err := c.getDiscussion(ctx, arg)
		if err != nil {
			return err
		}
		discussions[i] = d
	}
	encoder := json.NewEncoder(c.stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(&Discussions{
		Discussions: discussions,
	}); err != nil {
		return fmt.Errorf("output discussions as JSON: %w", err)
	}
	return nil
}

func (c *Controller) getDiscussion(ctx context.Context, arg string) (*Discussion, error) {
	pd, err := parseArg(arg)
	if err != nil {
		return nil, err
	}
	// Find and read a configuration file.
	// Get a discussion by GitHub GraphQL API.
	d, err := c.gh.GetDiscussion(ctx, pd.Owner, pd.Name, pd.Number)
	if err != nil {
		return nil, fmt.Errorf("get a discussion by GitHub API: %w", err)
	}
	discussion := convertDiscussion(d)
	discussion.Repo = &Repository{
		Owner: pd.Owner,
		Name:  pd.Name,
	}
	return discussion, nil
}

func convertLabels(in *github.Labels) []string {
	if in == nil {
		return nil
	}
	labels := make([]string, len(in.Nodes))
	for i, n := range in.Nodes {
		labels[i] = n.Name
	}
	return labels
}

type Discussions struct {
	Discussions []*Discussion
}

type Repository struct {
	Owner string
	Name  string
}

type Discussion struct {
	ID             string
	Title          string
	Body           string
	URL            string
	ClosedAt       string
	CreatedAt      string
	UpdatedAt      string
	AnswerChosenAt string
	UpvoteCount    int
	Repo           *Repository
	Author         *github.User
	Category       *github.Category
	Comments       []*Comment
	Labels         []string
	Answer         *Answer
	Reactions      map[string]*Reaction
	Poll           *Poll
	Locked         bool
	Closed         bool
}

func convertDiscussion(in *github.Discussion) *Discussion {
	return &Discussion{
		ID:             in.ID,
		Title:          in.Title,
		Body:           in.Body,
		URL:            in.URL,
		ClosedAt:       in.ClosedAt,
		CreatedAt:      in.CreatedAt,
		UpdatedAt:      in.UpdatedAt,
		AnswerChosenAt: in.AnswerChosenAt,
		UpvoteCount:    in.UpvoteCount,
		Author:         in.Author,
		Category:       in.Category,
		Comments:       convertComments(in.Comments),
		Labels:         convertLabels(in.Labels),
		Answer:         convertAnswer(in.Answer),
		Reactions:      convertReactions(in.Reactions),
		Poll:           convertPoll(in.Poll),
		Locked:         in.Locked,
		Closed:         in.Closed,
	}
}

type Poll struct {
	Question       string
	TotalVoteCount int
	Options        []*Option
}

func convertPoll(in *github.Poll) *Poll {
	if in == nil {
		return nil
	}
	return &Poll{
		Question:       in.Question,
		TotalVoteCount: in.TotalVoteCount,
		Options:        convertOptions(in.Options),
	}
}

func convertOptions(in *github.Options) []*Option {
	if in == nil {
		return nil
	}
	options := make([]*Option, len(in.Nodes))
	for i, n := range in.Nodes {
		options[i] = &Option{
			Option:         n.Option,
			TotalVoteCount: n.TotalVoteCount,
		}
	}
	return options
}

type Option struct {
	Option         string
	TotalVoteCount int
}

type Answer struct {
	ID          string
	Body        string
	CreatedAt   string
	UpvoteCount int
	Author      *github.User
	Reactions   map[string]*Reaction
}

func convertAnswer(in *github.Answer) *Answer {
	if in == nil {
		return nil
	}
	return &Answer{
		ID:          in.ID,
		Body:        in.Body,
		UpvoteCount: in.UpvoteCount,
		CreatedAt:   in.CreatedAt,
		Author:      in.Author,
		Reactions:   convertReactions(in.Reactions),
	}
}

func convertReactions(in *github.Reactions) map[string]*Reaction {
	if in == nil {
		return nil
	}
	m := map[string]*Reaction{}
	for _, n := range in.Nodes {
		emoji := convertReaction(n.Content)
		a, ok := m[emoji]
		if !ok {
			a = &Reaction{
				Emoji: emoji,
			}
		}
		a.Count++
		m[emoji] = a
	}
	return m
}

func convertReaction(s string) string {
	m := map[string]string{
		"THUMBS_UP":   "👍",
		"THUMBS_DOWN": "👎",
		"LAUGH":       "😄",
		"HOORAY":      "🎉",
		"CONFUSED":    "😕",
		"HEART":       "❤️",
		"ROCKET":      "🚀",
		"EYES":        "👀",
	}
	a, ok := m[s]
	if ok {
		return a
	}
	return s
}

type Reaction struct {
	Emoji string
	Count int
}

type Comment struct {
	ID              string
	Body            string
	URL             string
	CreatedAt       string
	MinimizedReason string
	UpvoteCount     int
	Author          *github.User
	Reactions       map[string]*Reaction
	Replies         []*Reply
	IsAnswer        bool
	IsMinimized     bool
}

func convertComments(in *github.Comments) []*Comment {
	if in == nil {
		return nil
	}
	comments := []*Comment{}
	for _, n := range in.Nodes {
		comments = append(comments, &Comment{
			ID:              n.ID,
			Body:            n.Body,
			URL:             n.URL,
			Author:          n.Author,
			CreatedAt:       n.CreatedAt,
			MinimizedReason: n.MinimizedReason,
			Reactions:       convertReactions(n.Reactions),
			Replies:         convertReplies(n.Replies),
			UpvoteCount:     n.UpvoteCount,
			IsAnswer:        n.IsAnswer,
			IsMinimized:     n.IsMinimized,
		})
	}
	return comments
}

type Reply struct {
	ID              string
	Body            string
	URL             string
	CreatedAt       string
	MinimizedReason string
	UpvoteCount     int
	Reactions       map[string]*Reaction
	Author          *github.User
	IsAnswer        bool
	IsMinimized     bool
}

func convertReplies(in *github.Replies) []*Reply {
	if in == nil {
		return nil
	}
	replies := make([]*Reply, len(in.Nodes))
	for i, n := range in.Nodes {
		replies[i] = &Reply{
			ID:              n.ID,
			Body:            n.Body,
			URL:             n.URL,
			MinimizedReason: n.MinimizedReason,
			UpvoteCount:     n.UpvoteCount,
			Reactions:       convertReactions(n.Reactions),
			Author:          n.Author,
			CreatedAt:       n.CreatedAt,
			IsAnswer:        n.IsAnswer,
			IsMinimized:     n.IsMinimized,
		}
	}
	return replies
}
