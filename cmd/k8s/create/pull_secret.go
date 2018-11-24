package create

import (
	"deployer/pkg"
	"deployer/pkg/deploy/k8s"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(pullSecretCmd)
}

var pullSecretCmd = &cobra.Command{
	Use:   "pull-secret",
	Short: "Create the docker registry pull secret",
	Run: func(cmd *cobra.Command, args []string) {
		registryDetails := pkg.GetDockerRegistryDetails()
		k8s.CreatePullSecret(registryDetails)
	},
}
