package delete

import (
	"deployer/pkg/deploy/k8s"
	"log"

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
		err := k8s.DeleteHelmServiceAccount(helmServiceUser)
		if err != nil {
			log.Fatal(err)
		}
	},
}
