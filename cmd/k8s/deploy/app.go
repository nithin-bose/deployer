package deploy

import (
	"deployer/pkg"
	"deployer/pkg/deploy"
	"deployer/pkg/deploy/k8s"
	"fmt"

	"github.com/spf13/cobra"
)

var ci bool

func init() {
	appCmd.Flags().BoolVarP(&ci, "ci", "c", false, "Indicate if command is run on the CI server")
	RootCmd.AddCommand(appCmd)
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Deploy service apps. Helm is required to be installed and configured",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			pkg.FatalF("Command must have exactly 3 arguments environment, app and version \n")
		}

		deploy.ValidateEnvironment(args[0])
		if args[0] == "production" && args[2] == "latest" {
			pkg.FatalF("Only versioned releases should be deployed to production \n")
		}
		fmt.Sprintf("Deploying %s... ", args[0])

		if ci {
			k8s.SetupKubeConfig(args[0])
		}
		k8s.ServiceApp(chartsDir, args[0], args[1], args[2])
	},
}
