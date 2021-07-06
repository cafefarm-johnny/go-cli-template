package cmd

import (
	"fmt"
	"go-cli-template/src/cmd/template"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:           "my-cli",
	Short:         "cobra cli template project.",
	Example:       "  my-cli [command] [-flags?]",
	SilenceUsage:  false, // usage 미출력 옵션 (에러 출력 시 usage가 함께 출력되므로 비활성화 처리)
	SilenceErrors: true,  // 에러 미출력 옵션 (RunE 함수에서 에러 반환 시 에러가 두번 출력되므로 비활성화 처리)
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`Usage: my-cli [command] if you want help? just enter "my-cli help"`)
	},
}

func Exec() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.SetHelpTemplate(template.HelpTemplate)

	rootCmd.Flags().BoolVarP(&help, "help", "h", false, "")
	rootCmd.Flags().Lookup("help").Hidden = true

	rootCmd.AddCommand(versionCmd)
}
