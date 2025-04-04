package cli

import (
	"context"
	"fmt"
	"io"

	"github.com/spf13/afero"
	"github.com/suzuki-shunsuke/ghd2i/pkg/controller"
	"github.com/urfave/cli/v3"
)

type createConfigCommand struct {
	stdout io.Writer
}

func (rc *createConfigCommand) command() *cli.Command {
	return &cli.Command{
		Name:  "create-config",
		Usage: "Create a configuration file",
		Description: `Create a configuration file.

$ ghd2i create-config
`,
		Action: rc.action,
	}
}

func (rc *createConfigCommand) action(_ context.Context, _ *cli.Command) error {
	ctrl, err := controller.New(rc.stdout, nil, afero.NewOsFs())
	if err != nil {
		return fmt.Errorf("initialize a controller: %w", err)
	}
	return ctrl.CreateConfig() //nolint:wrapcheck
}
