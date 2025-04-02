package cli

import (
	"context"
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghd2i/pkg/controller"
	"github.com/suzuki-shunsuke/ghd2i/pkg/github"
	"github.com/suzuki-shunsuke/ghd2i/pkg/log"
	"github.com/urfave/cli/v3"
)

type runCommand struct {
	logE   *logrus.Entry
	stdout io.Writer
}

func (rc *runCommand) command() *cli.Command { //nolint:funlen
	return &cli.Command{
		Name:  "run",
		Usage: "Create GitHub Issues from GitHub Discussions",
		Description: `Create GitHub Issues from GitHub Discussions

$ ghd2i run https://github.com/suzuki-shunsuke/test-github-action/discussions/55
`,
		Action: rc.action,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Aliases: []string{"c"},
				Usage:   "configuration file path. Configuration file is optional. If \\.ghd2i.yaml exists, it's used as the configuration file by default",
			},
			&cli.StringFlag{
				Name:  "data",
				Usage: "data file path. If data file path is set, the data is read from the file instead of calling GitHub API",
			},
			&cli.BoolFlag{
				Name:  "lock-discussion",
				Usage: "Lock discussions",
			},
			&cli.BoolFlag{
				Name:  "close-discussion",
				Usage: "Close discussions",
			},
			&cli.StringFlag{
				Name:  "lock",
				Usage: "Whether created issues are locked. One of 'auto', 'always', 'never'. Auto means that the issue is locked if the discussion is locked",
				Value: "auto",
			},
			&cli.StringFlag{
				Name:  "close",
				Usage: "Whether created issues are closed. One of 'auto', 'always', 'never'. Auto means that the issue is closed if the discussion is closed",
				Value: "auto",
			},
			&cli.StringFlag{
				Name:  "repo-owner",
				Usage: "Repository owner where issues are created. By default, issues are created in the repository of each discussion",
			},
			&cli.StringFlag{
				Name:  "repo-name",
				Usage: "Repository name where issues are created. By default, issues are created in the repository of each discussion",
			},
			&cli.StringFlag{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "A query to search discussions",
			},
			&cli.StringSliceFlag{
				Name:    "label",
				Aliases: []string{"l"},
				Usage:   "Additional labels to created issues",
			},
			&cli.StringSliceFlag{
				Name:    "assignee",
				Aliases: []string{"a"},
				Usage:   "Assignees of created issues",
			},
			&cli.BoolFlag{
				Name:  "dry-run",
				Usage: "Instead of creating issues, output issue body and comment bodies",
			},
		},
	}
}

func (rc *runCommand) action(ctx context.Context, c *cli.Command) error {
	logE := rc.logE
	log.SetLevel(c.String("log-level"), logE)
	log.SetColor(c.String("log-color"), logE)
	gh := github.New(ctx, os.Getenv("GITHUB_TOKEN"))
	ctrl, err := controller.New(rc.stdout, gh, afero.NewOsFs())
	if err != nil {
		return fmt.Errorf("initialize a controller: %w", err)
	}
	return ctrl.Run(ctx, logE, &controller.Param{ //nolint:wrapcheck
		ConfigFilePath:  c.String("config"),
		DataFilePath:    c.String("data"),
		Close:           c.String("close"),
		Lock:            c.String("lock"),
		RepoOwner:       c.String("repo-owner"),
		RepoName:        c.String("repo-name"),
		Query:           c.String("query"),
		Labels:          c.StringSlice("label"),
		Assignees:       c.StringSlice("assignee"),
		DryRun:          c.Bool("dry-run"),
		CloseDiscussion: c.Bool("close-discussion"),
		LockDiscussion:  c.Bool("lock-discussion"),
		Args:            c.Args().Slice(),
	})
}
