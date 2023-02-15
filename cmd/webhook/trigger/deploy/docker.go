package deploy

import (
	"deployer/pkg/trigger"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

var service string
var dockerStacksDir string

func init() {
	dockerCmd.PersistentFlags().StringVarP(&service, "service", "s", "", "Service in the app")
	dockerCmd.PersistentFlags().StringVarP(&dockerStacksDir, "docker-stacks-dir", "d", "", "Directory where docker compose app folders are present")
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
		err := trigger.DockerDeployApp(dockerStacksDir, args[0], service)
		if err != nil {
			log.Fatal(err)
		}
	},
}
