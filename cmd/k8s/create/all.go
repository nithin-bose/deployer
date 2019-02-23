package create

import (
	"deployer/pkg"
	"deployer/pkg/deploy/k8s"
	"log"
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
		err := k8s.SetRoleForDashboard()
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
		registryDetails := pkg.GetDockerRegistryDetails()
		err = k8s.CreatePullSecret(registryDetails)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
		err = k8s.CreateHelmServiceAccount(pkg.HelmServiceUser)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
		err = k8s.InstallHelm(pkg.HelmServiceUser)
		if err != nil {
			log.Fatal(err)
		}
	},
}
