package create

import (
	"deployer/pkg"
	"deployer/pkg/deploy/k8s"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(helmCmd)
}

var helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Install helm on cluster",
	Run: func(cmd *cobra.Command, args []string) {
		err := k8s.InstallHelm(pkg.HelmServiceUser)
		if err != nil {
			log.Fatal(err)
		}
	},
}
