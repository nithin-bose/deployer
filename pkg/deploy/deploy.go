package deploy

import "deployer/pkg"

func ValidateEnvironment(environment string) {
	if environment != "staging" && environment != "production" {
		pkg.FatalF("Unknown environment. Supported environments are staging and production \n")
	}
}
