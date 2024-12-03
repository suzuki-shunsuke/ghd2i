package controller

import (
	"errors"
	"fmt"

	"github.com/spf13/afero"
	"gopkg.in/yaml.v3"
)

func findConfig(fs afero.Fs, p string) (afero.File, error) {
	if p != "" {
		f, err := fs.Open(p)
		if err != nil {
			return nil, fmt.Errorf("open a configuration file: %w", err)
		}
		return f, nil
	}
	f, err := fs.Open("ghd2i.yaml")
	if err == nil {
		return f, nil
	}
	return fs.Open(".ghd2i.yaml") //nolint:wrapcheck
}

func findAndReadConfig(fs afero.Fs, cfg *Config, path string) error {
	f, err := findConfig(fs, path)
	if err != nil {
		if errors.Is(err, afero.ErrFileNotFound) {
			return nil
		}
		return fmt.Errorf("find a configuration file: %w", err)
	}
	defer f.Close()
	if err := yaml.NewDecoder(f).Decode(cfg); err != nil {
		return fmt.Errorf("decode a configuration file: %w", err)
	}
	return nil
}
