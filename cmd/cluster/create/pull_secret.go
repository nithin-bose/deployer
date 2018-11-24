package create

import (
	"deployer/pkg"
	"deployer/pkg/deploy"

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
		deploy.K8sCreatePullSecret(registryDetails)
	},
}
