package deploy

import (
	"deployer/pkg"
)

func ValidateCloudPlatforn(platform string) {
	if platform != "aws" && platform != "gce" {
		pkg.FatalF("Invalid cloud platform. Only aws and gce are supported  \n")
	}
}

func ValidateEnvironment(environment string) {
	if environment != "staging" && environment != "production" {
		pkg.FatalF("Unknown environment. Supported environments are staging and production \n")
	}
}
