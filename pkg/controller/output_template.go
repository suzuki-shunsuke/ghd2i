package controller

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

func (c *Controller) CreateConfig() error {
	cfg := &Config{
		IssueTemplate:   string(issueBodyTplByte),
		CommentTemplate: string(issueCommentBodyTplByte),
	}
	f, err := c.fs.Create("ghd2i.yaml")
	if err != nil {
		return fmt.Errorf("create a configuration file: %w", err)
	}
	defer f.Close()
	if err := yaml.NewEncoder(f).Encode(cfg); err != nil {
		return fmt.Errorf("write a configuration file: %w", err)
	}
	return nil
}
