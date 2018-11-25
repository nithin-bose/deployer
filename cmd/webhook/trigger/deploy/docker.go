package deploy

import (
	"deployer/pkg"
	"deployer/pkg/trigger"
	"fmt"

	"github.com/spf13/cobra"
)

var composeFile string
var composeFileDir string

func init() {
	dockerCmd.PersistentFlags().StringVarP(&composeFile, "compose-file", "f", "", "Path where the docker compose file is present")
	RootCmd.PersistentFlags().StringVarP(&composeFileDir, "compose-file-dir", "d", "", "Dir where docker-compose is to be run")
	RootCmd.AddCommand(dockerCmd)
}

var dockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Trigger docker deploy webhook",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			pkg.FatalF("Command must have exactly 1 argument, app name \n")
		}

		fmt.Sprintf("Deploying %s... ", args[0])
		trigger.DockerDeployApp(composeFileDir, composeFile, args[0])
	},
}
