package deploy

import (
	"github.com/spf13/cobra"
)

var chartsDir string

func init() {
	RootCmd.PersistentFlags().StringVarP(&chartsDir, "charts-dir", "d", "", "Path where the charts directory is present")
}

var RootCmd = &cobra.Command{
	Use:   "deploy",
	Short: "Deployment tasks. Helm is required to be installed and configured",
}
