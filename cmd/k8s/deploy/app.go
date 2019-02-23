package deploy

import (
	"deployer/pkg/deploy/k8s"
	"log"

	"github.com/spf13/cobra"
)

var ci bool
var force bool

func init() {
	appCmd.Flags().BoolVarP(&ci, "ci", "c", false, "Indicate if command is run on the CI server")
	appCmd.Flags().BoolVarP(&force, "force", "f", false, "Bypass versioned relase check for production")
	RootCmd.AddCommand(appCmd)
}

var appCmd = &cobra.Command{
	Use:   "app",
	Short: "Deploy service apps. Helm is required to be installed and configured",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 3 {
			log.Fatal("Command must have exactly 3 arguments environment, app and version \n")
		}
		err := k8s.DeployServiceApp(chartsDir, force, ci, args[0], args[1], args[2])
		if err != nil {
			log.Fatal("An error occurred:\n %s \n", err.Error())
		}
	},
}
