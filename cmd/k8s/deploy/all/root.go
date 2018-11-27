package all

import (
	"github.com/spf13/cobra"
)

var chartsDir string

func init() {
	RootCmd.PersistentFlags().StringVarP(&chartsDir, "charts-dir", "d", "", "Path where the charts directory is present")
}

var RootCmd = &cobra.Command{
	Use:   "all",
	Short: "Deploy all charts",
}
