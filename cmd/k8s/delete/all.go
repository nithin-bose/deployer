package delete

import (
	"deployer/pkg/deploy/k8s"
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
		k8s.DeleteHelmServiceAccount(helmServiceUser)
		time.Sleep(5 * time.Second)
		k8s.DeletePullSecret()
	},
}
