package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "grabit-cli",
	Short: "Grabit is a fantastic private message exchange platform",
	Long:  `Grabit is an exceptional private message exchange platform that offers fast and encrypted communication, with messages being automatically destroyed once read.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Check")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
