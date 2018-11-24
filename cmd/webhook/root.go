package webhook

import (
	"deployer/cmd/webhook/trigger"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(runCmd)
	RootCmd.AddCommand(trigger.RootCmd)
}

var RootCmd = &cobra.Command{
	Use:   "webhook",
	Short: "Webhook tasks",
}
