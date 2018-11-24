package k8s

import (
	"deployer/pkg"
)

func ValidateCloudPlatform(platform string) {
	if platform != "aws" && platform != "gce" {
		pkg.FatalF("Invalid cloud platform. Only aws and gce are supported  \n")
	}
}
