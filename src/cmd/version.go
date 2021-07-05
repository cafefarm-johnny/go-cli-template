package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"go-cli-template/src/app"
)

var versionCmd = &cobra.Command{
	Use:     "version",
	Short:   "  print the application version.",
	Example: "my-cli version",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(fmt.Sprintf("%s (ver %s)", app.Name, app.Version))
	},
}
