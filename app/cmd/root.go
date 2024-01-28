package cmd

import (
	"fmt"
	"grabit-cli/app/cmd/term"
	common_adapters "grabit-cli/core/common/gateways/adapters"
	id_adapters "grabit-cli/core/identities/gateways/adapters"
	"grabit-cli/core/identities/usecases"
	"os"

	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

var display = term.PTermDisplay{}
var ptermLogger = pterm.DefaultLogger.WithLevel(pterm.LogLevelTrace)
var logger = &common_adapters.PtermLogger{Logger: *ptermLogger}

var iddentityGateway = id_adapters.NewFsIdentityGateway(id_adapters.IdentityPath{Dir: ".grabit", FileName: ".identity"}, "identity.pub", "private.pem", logger)
var rootCmd = &cobra.Command{
	Use:   "grabit-cli",
	Short: "Grabit is a fantastic private message exchange platform",
	Long:  `Grabit is an exceptional private message exchange platform that offers fast and encrypted communication, with messages being automatically destroyed once read.`,
	Run: func(cmd *cobra.Command, args []string) {
		display.DefaultText("Send private messages anonymously")
		display.Title("grabit")
		display.Info("Checking Grabit identity....")
		fsInstallation.
		useCase := usecases.NewGetCurrentIdentityUseCase(&iddentityGateway)
		response := useCase.Execute(usecases.GetCurrentIdentityRequest{Email: })
		if response != nil {
			pterm.Success.Println("We found Grabit identity ", pterm.Gray(response.Email), "You are ready to send and receive secure notes !")
			return
		}
		display.Error("We cannot found your Grabit identity, you should create a new one using commande line:", "")
		pterm.DefaultBasicText.Println(pterm.Gray("$ grabit-cli init"))

	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
