package webhook

import (
	"deployer/pkg/webhook"

	"github.com/spf13/cobra"
)

func init() {
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a webserver that exposes the CLI options as webhooks",
	Run: func(cmd *cobra.Command, args []string) {
		webhook.Run()
	},
}
