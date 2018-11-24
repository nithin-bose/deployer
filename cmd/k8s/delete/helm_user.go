package delete

import (
	"deployer/pkg/deploy/k8s"

	"github.com/spf13/cobra"
)

var helmServiceUser string

func init() {
	helmServiceUser = "helm"
	RootCmd.AddCommand(helmUserCmd)
}

var helmUserCmd = &cobra.Command{
	Use:   "helm-user",
	Short: "Delete the service account created for helm to use",
	Run: func(cmd *cobra.Command, args []string) {
		k8s.DeleteHelmServiceAccount(helmServiceUser)
	},
}
