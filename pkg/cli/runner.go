package cli

import (
	"context"
	"io"

	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v3"
)

type Runner struct {
	Stdin   io.Reader
	Stdout  io.Writer
	Stderr  io.Writer
	LDFlags *LDFlags
	LogE    *logrus.Entry
}

type LDFlags struct {
	Version string
	Commit  string
	Date    string
}

func (r *Runner) Run(ctx context.Context, args ...string) error {
	app := cli.Command{
		Name:    "ghd2i",
		Usage:   "Create GitHub Issues from GitHub Discussions",
		Version: r.LDFlags.Version + " (" + r.LDFlags.Commit + ")",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "log-level",
				Usage: "log level",
			},
			&cli.StringFlag{
				Name:  "log-color",
				Usage: "Log color. One of 'auto' (default), 'always', 'never'",
			},
		},
		EnableShellCompletion: true,
		Commands: []*cli.Command{
			(&versionCommand{
				stdout:  r.Stdout,
				version: r.LDFlags.Version,
				commit:  r.LDFlags.Commit,
			}).command(),
			(&runCommand{
				logE:   r.LogE,
				stdout: r.Stdout,
			}).command(),
			(&createConfigCommand{
				stdout: r.Stdout,
			}).command(),
			(&getDiscussionCommand{
				logE:   r.LogE,
				stdout: r.Stdout,
			}).command(),
			(&completionCommand{
				logE:   r.LogE,
				stdout: r.Stdout,
			}).command(),
		},
	}

	return app.Run(ctx, args) //nolint:wrapcheck
}
