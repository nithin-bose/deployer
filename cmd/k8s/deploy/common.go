package deploy

import (
	"deployer/pkg/deploy/k8s"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(commonCmd)
}

var commonCmd = &cobra.Command{
	Use:   "common",
	Short: "Deploy common apps",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("Command must have exactly 2 arguments, environment and app.  \n")
		}

		err := k8s.CommonApp(chartsDir, args[1], args[0])
		if err != nil {
			log.Fatal(err)
		}
	},
}
