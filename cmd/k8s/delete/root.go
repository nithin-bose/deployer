package delete

import (
	"github.com/spf13/cobra"
)

func init() {
}

var RootCmd = &cobra.Command{
	Use:   "delete",
	Short: "Delete cluster resources",
}
