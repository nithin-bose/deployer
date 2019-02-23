package deploy

import (
	"deployer/pkg/trigger"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(k8sCmd)
}

var k8sCmd = &cobra.Command{
	Use:   "k8s",
	Short: "Trigger k8s deploy webhook",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			log.Fatal("Command must have exactly 3 arguments environment, app and version")
		}
		err := trigger.K8sDeployApp(args[0], args[1], args[2])
		if err != nil {
			log.Fatal(err)
		}
	},
}
