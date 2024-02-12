package gateways

type Logger interface {
	Trace(message string)
	Info(message string)
	Error(message string, err error)
}
