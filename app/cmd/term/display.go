package term

type Display interface {
	Title(title string)
	Info(message string, format string)
	Error(message string, format string)
	Success(message string, format string)
	CenteredText(message string, format string)
	Text(message string)
	DefaultText(message string)
}
