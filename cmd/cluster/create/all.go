package create

import (
	"deployer/pkg"
	"deployer/pkg/deploy"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(allCmd)
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Create all cluster resources",
	Run: func(cmd *cobra.Command, args []string) {
		deploy.K8sSetRoleForDashboard()
		time.Sleep(5 * time.Second)
		registryDetails := pkg.GetDockerRegistryDetails()
		deploy.K8sCreatePullSecret(registryDetails)
		time.Sleep(5 * time.Second)
		deploy.K8sCreateHelmServiceAccount(pkg.HelmServiceUser)
		time.Sleep(5 * time.Second)
		deploy.K8sInstallHelm(pkg.HelmServiceUser)
	},
}
