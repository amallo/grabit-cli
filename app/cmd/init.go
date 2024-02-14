package cmd

import (
	"github.com/pterm/pterm"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(initCmd)
}

var initCmd = &cobra.Command{
	Use:   "init",
	Short: "Create you identity",
	Run: func(cmd *cobra.Command, args []string) {
		pterm.Info.Println("Creating your identity....")
		/*
			nameInput, _ := pterm.DefaultInteractiveTextInput.Show("What is your name")
			getCurrentIdentity := usecases.NewGetCurrentIdentityUseCase(&iddentityGateway)
			identityResponse := getCurrentIdentity.Execute(usecases.GetCurrentIdentityRequest{Name: nameInput})

			emailInput, _ := pterm.DefaultInteractiveTextInput.Show("What is your email")

			if identityResponse != nil && identityResponse.Email == emailInput {
				pterm.Success.Println("Account", emailInput, "already exists.", "\nYou can now send message using: ", pterm.Gray("$ grabit-cli send"))
				return
			}

			passwordInputMask := pterm.DefaultInteractiveTextInput.WithMask("*")
			passwordInput1, _ := passwordInputMask.Show("Enter your password")
			passwordInput2, _ := passwordInputMask.Show("Enter your password again")
			if passwordInput1 != passwordInput2 {
				pterm.Error.Println("Sorry passwords not match")
				return
			}

			registerIdentity := usecases.NewRegisterIdentityUseCase(&iddentityGateway)
			err := registerIdentity.Execute(usecases.RegisterIdentityParams{Email: emailInput, Name: emailInput, PassPhrase: passwordInput1})
			if err != nil && err.Error() == "ALREADY_HAVE_IDENTITY" {
				pterm.Error.Println("Account", emailInput, "was already created.")
				return
			}
			pterm.Success.Printfln("Congratulations, you are now ready to share messages with your buddies !")*/
	},
}
