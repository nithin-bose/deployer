package create

import (
	"deployer/pkg"
	"deployer/pkg/deploy/k8s"
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
		k8s.SetRoleForDashboard()
		time.Sleep(5 * time.Second)
		registryDetails := pkg.GetDockerRegistryDetails()
		k8s.CreatePullSecret(registryDetails)
		time.Sleep(5 * time.Second)
		k8s.CreateHelmServiceAccount(pkg.HelmServiceUser)
		time.Sleep(5 * time.Second)
		k8s.InstallHelm(pkg.HelmServiceUser)
	},
}
