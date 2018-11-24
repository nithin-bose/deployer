package create

import (
	"deployer/pkg"
	"deployer/pkg/deploy/k8s"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(helmUserCmd)
}

var helmUserCmd = &cobra.Command{
	Use:   "helm-user",
	Short: "Create a service account for helm to use",
	Run: func(cmd *cobra.Command, args []string) {
		k8s.CreateHelmServiceAccount(pkg.HelmServiceUser)
	},
}
