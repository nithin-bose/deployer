package deploy

import (
	"deployer/pkg"
	"deployer/pkg/deploy/docker"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(appCmd)
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Deploy service apps. docker-compose is required to be installed and configured",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			pkg.FatalF("Command must have exactly 1 argument, app name \n")
		}

		fmt.Sprintf("Deploying %s... ", args[0])
		docker.DeployServiceApp(composeFile, args[0])
	},
}
