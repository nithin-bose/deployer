package all

import (
	"bufio"
	"deployer/pkg/export/s3"
	"fmt"
	"io"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	RootCmd.AddCommand(dockerImagesCmd)
}

var dockerImagesCmd = &cobra.Command{
	Use:   "docker-images",
	Short: "Export all docker images specified in a file to s3",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			log.Fatal("Command must have exactly 1 argument, path to a file with a list of docker image repositories \n")
		}

		f, err := os.Open(args[0])
		if err != nil {
			log.Fatalf("An error occurred:\n %s \n", err.Error())
		}
		defer f.Close()

		buf := bufio.NewReader(f)
		for {
			l, _, err := buf.ReadLine()

			// If we're just at the EOF, break
			if err != nil {
				if err == io.EOF {
					break
				} else {
					log.Fatalf("An error occurred:\n %s \n", err.Error())
				}
			}
			line := strings.TrimSpace(string(l))
			split := strings.Split(line, ":")
			uploadDetails, err := s3.ExportDockerImageToS3(split[0], split[1])
			if err != nil {
				log.Fatalf("An error occurred:\n %s \n", err.Error())
			}
			fmt.Printf("%s successfully uploaded to %s \n", line, uploadDetails.Location)
		}
	},
}
