package all

import (
	"github.com/spf13/cobra"
)

func init() {
}

var RootCmd = &cobra.Command{
	Use:   "all",
	Short: "Export all to s3",
}
