package cli

import (
	"io"

	"github.com/suzuki-shunsuke/ghd2i/pkg/controller"
	"github.com/urfave/cli/v2"
)

type outputTemplateCommand struct {
	stdout io.Writer
}

func (rc *outputTemplateCommand) command() *cli.Command {
	return &cli.Command{
		Name:  "output-template",
		Usage: "Output the default templates",
		Description: `Get the default templates.

$ ghd2i output-template
`,
		Action: rc.action,
	}
}

func (rc *outputTemplateCommand) action(c *cli.Context) error {
	ctrl, err := controller.New(rc.stdout, nil, nil)
	if err != nil {
		return err
	}
	return ctrl.OutputTemplate()
}
