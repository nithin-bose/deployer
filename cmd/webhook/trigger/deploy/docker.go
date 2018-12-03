package deploy

import (
	"deployer/pkg"
	"deployer/pkg/trigger"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
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
		trigger.DockerDeployApp(args[0])
	},
}
