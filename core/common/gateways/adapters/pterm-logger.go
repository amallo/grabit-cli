package adapters

import "github.com/pterm/pterm"

type PtermLogger struct {
	Logger pterm.Logger
}

func (l *PtermLogger) Trace(message string) {
	l.Logger.Trace(message)
}
func (l *PtermLogger) Info(message string) {
	l.Logger.Info(message)
}
func (l *PtermLogger) Error(message string, err error) {
	l.Logger.Error(message, l.Logger.Args("error", err.Error()))
}
