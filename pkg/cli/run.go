package cli

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghd2i/pkg/controller"
	"github.com/suzuki-shunsuke/ghd2i/pkg/github"
	"github.com/suzuki-shunsuke/ghd2i/pkg/log"
	"github.com/urfave/cli/v2"
)

type runCommand struct {
	logE   *logrus.Entry
	stdout io.Writer
}

func (rc *runCommand) command() *cli.Command {
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
				Usage:   "configuration file path",
			},
			&cli.StringFlag{
				Name:  "data",
				Usage: "data file path",
			},
			&cli.BoolFlag{
				Name:  "dry-run",
				Usage: "dry run",
			},
		},
	}
}

func (rc *runCommand) action(c *cli.Context) error {
	logE := rc.logE
	log.SetLevel(c.String("log-level"), logE)
	log.SetColor(c.String("log-color"), logE)
	gh := github.New(c.Context, os.Getenv("GITHUB_TOKEN"))
	ctrl, err := controller.New(rc.stdout, gh, afero.NewOsFs())
	if err != nil {
		return err
	}
	return ctrl.Run(c.Context, logE, &controller.Param{ //nolint:wrapcheck
		ConfigFilePath: c.String("config"),
		DataFilePath:   c.String("data"),
		DryRun:         c.Bool("dry-run"),
		Args:           c.Args().Slice(),
	})
}
