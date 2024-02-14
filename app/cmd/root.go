package cmd

import (
	"fmt"
	"os"

	"github.com/pterm/pterm"
	"github.com/pterm/pterm/putils"
	"github.com/spf13/cobra"
)

//var ptermLogger = pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
//var logger = &common_adapters.PtermLogger{Logger: *ptermLogger, Context: common_adapters.Context{Name: "root"}}

// var iddentityGateway = id_adapters.NewFsIdentityGateway(id_adapters.IdentityPath{Dir: ".grabit", FileName: ".identity"}, "identity.pub", "private.pem", logger)
var rootCmd = &cobra.Command{
	Use:   "grabit-cli",
	Short: "Grabit is a fantastic private message exchange platform",
	Long:  `Grabit is an exceptional private message exchange platform that offers fast and encrypted communication, with messages being automatically destroyed once read.`,
	Run: func(cmd *cobra.Command, args []string) {

		var letters = putils.LettersFromString("grabit")
		pterm.DefaultBigText.WithLetters(letters).Render()
		pterm.DefaultBasicText.Println(pterm.Gray("Send private messages anonymously"))

		/*
			f /*sInstallation.
			useCase := usecases.NewGetCurrentIdentityUseCase(&iddentityGateway)
			response := useCase.Execute(usecases.GetCurrentIdentityRequest{Email: })
			if response != nil {
				pterm.Success.Println("We found Grabit identity ", pterm.Gray(response.Email), "You are ready to send and receive secure notes !")
				return
			}
			display.Error("We cannot found your Grabit identity, you should create a new one using commande line:", "")
			pterm.DefaultBasicText.Println(pterm.Gray("$ grabit-cli init"))*/

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
