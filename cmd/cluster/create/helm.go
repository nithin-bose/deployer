package create

import (
	"deployer/pkg"
	"deployer/pkg/deploy"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(helmCmd)
}

var helmCmd = &cobra.Command{
	Use:   "helm",
	Short: "Install helm on cluster",
	Run: func(cmd *cobra.Command, args []string) {
		deploy.K8sInstallHelm(pkg.HelmServiceUser)
	},
}
