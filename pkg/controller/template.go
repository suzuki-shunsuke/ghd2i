package controller

import "text/template"

func parseTemplate(s string) (*template.Template, error) {
	return template.New("_").Parse(s)
}
