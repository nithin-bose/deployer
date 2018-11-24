package deploy

import (
	"deployer/pkg"
	"deployer/pkg/deploy"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(infraCmd)
}

var infraCmd = &cobra.Command{
	Use:   "infra",
	Short: "Deploy infrastructure apps like traefik, volume provisioners etc. Helm is required to be installed and configured",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			pkg.FatalF("Command must have exactly 2 arguments, cloud platform and app.  \n")
		}

		deploy.ValidateCloudPlatforn(args[0])
		deploy.InfraApp(chartsDir, args[0], args[1])
	},
}
