package s3

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	s3svc "github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

func Download(bucket string, key string, filePath string) error {
	// Create a file to write the S3 Object contents to.
	downloader := s3manager.NewDownloader(awsSession)

	f, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer f.Close()

	// Write the contents of S3 Object to the file
	_, err = downloader.Download(f, &s3svc.GetObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
	})
	if err != nil {
		return err
	}
	return nil
}
