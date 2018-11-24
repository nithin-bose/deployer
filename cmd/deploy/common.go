package deploy

import (
	"deployer/pkg"
	"deployer/pkg/deploy"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(systemCmd)
}

var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "Deploy apps to kube-system namespace",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			pkg.FatalF("Command must have exactly 2 arguments, environment and app.  \n")
		}
		deploy.ValidateEnvironment(args[0])
		deploy.SystemApp(chartsDir, args[1], args[0])
	},
}
