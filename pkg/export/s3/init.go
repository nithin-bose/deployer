package s3

import (
	"errors"
	"log"
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
)

var awsSession *session.Session

func Init() error {
	awsConfig := aws.Config{}

	minioEndpoint := os.Getenv("MINIO_ENDPOINT")
	if minioEndpoint != "" {
		log.Println("Detected minio server:", minioEndpoint)
		awsConfig.Endpoint = aws.String(minioEndpoint)
		awsConfig.S3ForcePathStyle = aws.Bool(true)
	} else {
		awsRegion := os.Getenv("AWS_REGION")
		if awsRegion == "" {
			return errors.New("AWS_REGION not set")
		}
		awsConfig.Region = aws.String(awsRegion)
	}

	awsSession = session.Must(session.NewSession(&awsConfig))
	return nil
}
