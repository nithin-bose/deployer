package deploy

import (
	"deployer/pkg/deploy/k8s"
	"log"

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
			log.Fatal("Command must have exactly 1 argument, app")
		}

		err := k8s.AdminPanel(chartsDir, args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}
