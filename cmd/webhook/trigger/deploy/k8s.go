package deploy

import (
	"deployer/pkg"
	"deployer/pkg/trigger"

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
			pkg.FatalF("Command must have exactly 3 arguments environment, app and version \n")
		}
		trigger.K8sDeployApp(args[0], args[1], args[2])
	},
}
