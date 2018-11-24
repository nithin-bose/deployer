package deploy

import (
	"deployer/pkg"
	"deployer/pkg/deploy/k8s"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(adminCmd)
}

var adminCmd = &cobra.Command{
	Use:   "admin",
	Short: "Deploy admin panels. Helm is required to be installed and configured",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			pkg.FatalF("Command must have exactly 1 argument, app.  \n")
		}

		k8s.AdminPanel(chartsDir, args[0])
	},
}
