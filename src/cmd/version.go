package cmd

import (
	"fmt"
	"go-cli-template/src/app"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "print the application version.",
	Example: "  my-cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s (ver %s)", app.Name, app.Version))
	},
}
