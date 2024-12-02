package controller

import (
	"fmt"
)

func (c *Controller) OutputTemplate() error {
	fmt.Fprintln(c.stdout, string(issueBodyTplByte))
	return nil
}
