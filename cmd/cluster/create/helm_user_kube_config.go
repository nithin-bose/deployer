package create

import (
	"deployer/pkg"
	"deployer/pkg/deploy"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(helmUserKubeConfigCmd)
}

var helmUserKubeConfigCmd = &cobra.Command{
	Use:   "helm-user-kube-config",
	Short: "Create the kube config for the helm service account to use with CI/CD",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			pkg.FatalF("Command must have exactly 1 argument, cluster name.  \n")
		}
		deploy.K8sCreateSAKubeConfig(pkg.HelmServiceUser, args[0])
	},
}
