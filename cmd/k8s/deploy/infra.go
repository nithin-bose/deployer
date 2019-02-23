package deploy

import (
	"deployer/pkg/deploy/k8s"
	"log"

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
			log.Fatal("Command must have exactly 2 arguments, cloud platform and app")
		}

		err := k8s.ValidateCloudPlatform(args[0])
		if err != nil {
			log.Fatal(err)
		}
		err = k8s.InfraApp(chartsDir, args[0], args[1])
		if err != nil {
			log.Fatal(err)
		}
	},
}
