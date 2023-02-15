package deploy

import (
	"deployer/pkg/deploy/docker"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(appCmd)
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Deploy apps. docker compose v2 is required to be installed and configured",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Command must have exactly 1 argument, app name")
		}

		fmt.Sprintf("Deploying %s... ", args[0])

		err := docker.DeployServiceApp(dockerStacksDir, composeFile, args[0], service)
		if err != nil {
			log.Fatal(err)
		}
	},
}
