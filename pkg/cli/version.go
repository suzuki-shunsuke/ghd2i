package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"io"

	"github.com/urfave/cli/v3"
)

type versionCommand struct {
	stdout  io.Writer
	version string
	commit  string
}

func (vc *versionCommand) command() *cli.Command {
	return &cli.Command{
		Name:   "version",
		Usage:  "Show version",
		Action: vc.action,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name: "json",
			},
		},
	}
}

func (vc *versionCommand) action(_ context.Context, c *cli.Command) error {
	if !c.Bool("json") {
		cli.ShowVersion(c)
		return nil
	}
	encoder := json.NewEncoder(vc.stdout)
	encoder.SetIndent("", "  ")
	if err := encoder.Encode(map[string]string{
		"version": vc.version,
		"commit":  vc.commit,
	}); err != nil {
		return fmt.Errorf("encode the version as JSON: %w", err)
	}
	return nil
}
