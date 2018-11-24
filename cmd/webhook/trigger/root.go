package trigger

import (
	"deployer/cmd/webhook/trigger/deploy"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(deploy.RootCmd)
}

var RootCmd = &cobra.Command{
	Use:   "trigger",
	Short: "Trigger deployer web hooks",
}
