package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var rootCmd = &cobra.Command{
	Use:           "my-cli",
	Short:         "cobra cli template project.",
	Example:       "my-cli [command] [-flags?]",
	SilenceErrors: true,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("  cobra cli tmepate project. if you want help my-cli help")
	},
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
}
