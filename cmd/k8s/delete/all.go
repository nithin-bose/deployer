package delete

import (
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
	Short: "Delete all cluster resources",
	Run: func(cmd *cobra.Command, args []string) {
		err := k8s.DeleteHelmServiceAccount(helmServiceUser)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(5 * time.Second)
		err = k8s.DeletePullSecret()
		if err != nil {
			log.Fatal(err)
		}
	},
}
