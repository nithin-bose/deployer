package webhook

import (
	"deployer/pkg/webhook"

	"github.com/spf13/cobra"
)

var socket bool

func init() {
	runCmd.Flags().BoolVarP(&socket, "socket", "s", false, "Webserver listens on /var/run/deployer.sock")
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Runs a webserver that exposes the CLI options as webhooks",
	Run: func(cmd *cobra.Command, args []string) {
		webhook.Run(socket)
	},
}
