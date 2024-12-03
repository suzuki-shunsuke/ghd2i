package cli

import (
	"fmt"
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/suzuki-shunsuke/ghd2i/pkg/controller"
	"github.com/suzuki-shunsuke/ghd2i/pkg/github"
	"github.com/urfave/cli/v2"
)

type getDiscussionCommand struct {
	stdout io.Writer
	logE   *logrus.Entry
}

func (rc *getDiscussionCommand) command() *cli.Command {
	return &cli.Command{
		Name:  "get-discussion",
		Usage: "Get discussion and output the data",
		Description: `Get discussion and output the data

$ ghd2i get-discussion <discussion-url> [<discussion-url> ...]
`,
		Action: rc.action,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "query",
				Aliases: []string{"q"},
				Usage:   "A query to search discussions. 'is:discussions' is added to the query",
			},
		},
	}
}

func (rc *getDiscussionCommand) action(c *cli.Context) error {
	gh := github.New(c.Context, os.Getenv("GITHUB_TOKEN"))
	ctrl, err := controller.New(rc.stdout, gh, nil)
	if err != nil {
		return fmt.Errorf("initialize a controller: %w", err)
	}
	return ctrl.GetDiscussion(c.Context, rc.logE, &controller.Param{ //nolint:wrapcheck
		Args:  c.Args().Slice(),
		Query: c.String("query"),
	})
}
