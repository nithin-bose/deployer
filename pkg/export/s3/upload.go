package s3

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
)

type UploadDetails struct {
	Bucket   string `json:"bucket"`
	Key      string `json:"key"`
	Location string `json:"location"`
}

func Upload(filePath string) (*UploadDetails, error) {
	// Create an uploader with the session and default options
	uploader := s3manager.NewUploader(awsSession)

	f, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	// Upload the file to S3.
	bucket := os.Getenv("DEPLOYER_EXPORT_IMAGE_S3_BUCKET")
	if bucket == "" {
		return nil, errors.New("DEPLOYER_EXPORT_IMAGE_S3_BUCKET not set")
	}

	key := filepath.Base(filePath)
	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   f,
	})
	if err != nil {
		return nil, err
	}

	uploadDetails := &UploadDetails{}
	uploadDetails.Bucket = bucket
	uploadDetails.Key = key
	uploadDetails.Location = aws.StringValue(&result.Location)
	return uploadDetails, nil
}
