package deploy

import (
	"github.com/spf13/cobra"
)

var composeFile string
var dockerStacksDir string
var service string

func init() {
	RootCmd.PersistentFlags().StringVarP(&composeFile, "compose-file", "f", "", "Path where the docker compose file is present")
	RootCmd.PersistentFlags().StringVarP(&dockerStacksDir, "docker-stacks-dir", "d", ".", "Directory where docker compose app folders are present")
	RootCmd.PersistentFlags().StringVarP(&service, "service", "s", "", "Service in the app")
}

var RootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deployment tasks. Docker compose is required to be installed and configured",
}
