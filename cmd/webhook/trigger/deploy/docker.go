package deploy

import (
	"deployer/pkg/trigger"
	"fmt"
	"log"

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
			log.Fatal("Command must have exactly 1 argument, app name")
		}

		fmt.Sprintf("Deploying %s... ", args[0])
		err := trigger.DockerDeployApp(args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}
