package grifts

import (
	"github.com/Noah-Huppert/gh-gantt/server/actions"
	"github.com/gobuffalo/buffalo"
)

func init() {
	buffalo.Grifts(actions.App())
}
