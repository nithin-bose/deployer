package s3

import (
	"deployer/pkg/export/s3"
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(dockerImageCmd)
}

var dockerImageCmd = &cobra.Command{
	Use:   "docker-image",
	Short: "Export docker image to s3",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			log.Fatal("Command must have exactly 2 arguments docker registry and tag \n")
		}
		uploadDetails, err := s3.ExportDockerImageToS3(args[0], args[1])
		if err != nil {
			log.Fatal("An error occurred:\n %s \n", err.Error())
		}
		fmt.Println("Successfully uploaded to", uploadDetails.Location)
	},
}
