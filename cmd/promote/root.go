package promote

import (
	"deployer/pkg/release"
	"log"

	"github.com/spf13/cobra"
)

func init() {
}

var RootCmd = &cobra.Command{
	Use:   "promote",
	Short: "Promote for release. Command must have exactly 4 arguments, source branch, target branch, projectID and user ID.",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 4 && len(args) > 5 {
			log.Fatal("Command must have atleast 4 arguments, source branch, target branch, projectID and user ID. Optional promoter name \n")
		}

		if len(args) == 4 {
			args[4] = ""
		}
		err := release.Promote(args[0], args[1], args[2], args[3], args[4])
		if err != nil {
			log.Fatal(err)
		}
	},
}
