package term

import (
	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
)

type PTermDisplay struct {
}

func (display *PTermDisplay) Title(text string) {
	var letters = putils.LettersFromString(text)
	pterm.DefaultBigText.WithLetters(letters).Render()
}

func (display *PTermDisplay) Info(message string, format string) {
	pterm.Info.Println(message, format)
}
func (display *PTermDisplay) Error(message string, format string) {
	pterm.Error.Println(message, format)
}

func (display *PTermDisplay) Success(message string, format string) {
	pterm.Success.Println(message, format)
}

func (display *PTermDisplay) CenteredText(message string, format string) {
	pterm.DefaultCenter.Println(pterm.Gray(message), format)
}
func (display *PTermDisplay) DefaultText(message string) {
	pterm.DefaultBasicText.Println(pterm.Gray(message))
}
func (display *PTermDisplay) Text(message string) {
	style := pterm.NewStyle(pterm.BgGray, pterm.FgLightBlue, pterm.Bold)
	pterm.DefaultBasicText.WithStyle(style).Printfln(message)
}
