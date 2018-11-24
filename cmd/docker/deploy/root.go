package deploy

import (
	"github.com/spf13/cobra"
)

var composeFile string

func init() {
	RootCmd.PersistentFlags().StringVarP(&composeFile, "compose-file", "f", "", "Path where the docker compose file is present")
}

var RootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deployment tasks. Docker compose is required to be installed and configured",
}
