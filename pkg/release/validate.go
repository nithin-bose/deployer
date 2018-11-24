package release

import (
	"deployer/pkg"
)

func ValidateType(release string) {
	if release != "major" && release != "minor" && release != "patch" {
		pkg.FatalF("Invalid release type. Only major, minor and patch are supported  \n")
	}
}
