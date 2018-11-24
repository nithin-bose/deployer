package delete

import (
	"deployer/pkg/deploy"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(allCmd)
}

var allCmd = &cobra.Command{
	Use:   "all",
	Short: "Delete all cluster resources",
	Run: func(cmd *cobra.Command, args []string) {
		deploy.K8sDeleteHelmServiceAccount(helmServiceUser)
		time.Sleep(5 * time.Second)
		deploy.K8sDeletePullSecret()
	},
}
