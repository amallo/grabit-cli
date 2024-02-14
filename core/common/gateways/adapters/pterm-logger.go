package adapters

import (
	"fmt"

	"github.com/pterm/pterm"
)

type Context struct {
	Name string
}
type PtermLogger struct {
	Logger  pterm.Logger
	Context Context
}

func (l PtermLogger) Trace(message string) {
	msg := fmt.Sprintf("[%s] %s", l.Context.Name, message)
	l.Logger.Trace(msg)
}
func (l PtermLogger) Info(message string) {
	msg := fmt.Sprintf("[%s] %s", l.Context.Name, message)
	l.Logger.Info(msg)
}
func (l PtermLogger) Error(message string, err error) {
	msg := fmt.Sprintf("[%s] %s", l.Context.Name, message)
	l.Logger.Error(msg, l.Logger.Args("error", err.Error()))
}
