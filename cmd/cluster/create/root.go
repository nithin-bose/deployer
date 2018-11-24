package create

import (
	"github.com/spf13/cobra"
)

func init() {
}

var RootCmd = &cobra.Command{
	Use:   "create",
	Short: "Create cluster resources",
}
