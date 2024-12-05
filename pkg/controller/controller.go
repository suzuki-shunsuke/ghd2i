package controller

import (
	"io"
	"text/template"

	"github.com/spf13/afero"
)

type Controller struct {
	stdout            io.Writer
	gh                GitHub
	fs                afero.Fs
	issueBody         *template.Template
	issueCommentBody  *template.Template
	title             *template.Template
	discussionComment *template.Template
}

func New(stdout io.Writer, gh GitHub, fs afero.Fs) (*Controller, error) {
	issueBodyTpl, err := parseTemplate(string(issueBodyTplByte))
	if err != nil {
		return nil, err
	}
	issueCommentBodyTpl, err := parseTemplate(string(issueCommentBodyTplByte))
	if err != nil {
		return nil, err
	}
	return &Controller{
		stdout:           stdout,
		gh:               gh,
		fs:               fs,
		issueBody:        issueBodyTpl,
		issueCommentBody: issueCommentBodyTpl,
	}, nil
}
