package cli

import (
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
	}
}

func (rc *getDiscussionCommand) action(c *cli.Context) error {
	gh := github.New(c.Context, os.Getenv("GITHUB_TOKEN"))
	ctrl, err := controller.New(rc.stdout, gh, nil)
	if err != nil {
		return err
	}
	return ctrl.GetDiscussion(c.Context, rc.logE, &controller.Param{
		Args: c.Args().Slice(),
	})
}
