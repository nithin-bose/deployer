package webhook

import (
	"deployer/pkg/webhook"

	"github.com/spf13/cobra"
)

func init() {
}

var RootCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Runs a webserver that exposes the CLI options as webhooks",
	Run: func(cmd *cobra.Command, args []string) {
		webhook.Run()
	},
}
