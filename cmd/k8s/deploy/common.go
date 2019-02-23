package deploy

import (
	"deployer/pkg/deploy"
	"deployer/pkg/deploy/k8s"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(systemCmd)
}

var systemCmd = &cobra.Command{
	Use:   "system",
	Short: "Deploy apps to kube-system namespace",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("Command must have exactly 2 arguments, environment and app.  \n")
		}
		err := deploy.ValidateEnvironment(args[0])
		if err != nil {
			log.Fatal(err)
		}
		err = k8s.SystemApp(chartsDir, args[1], args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}
