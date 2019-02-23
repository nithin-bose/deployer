package k8s

import "errors"

func ValidateCloudPlatform(platform string) error {
	if platform != "aws" && platform != "gce" {
		return errors.New("Invalid cloud platform. Only aws and gce are supported")
	}
	return nil
}
