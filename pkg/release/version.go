package release

import (
	"deployer/pkg"
	"fmt"
	"strconv"
	"strings"
)

func newVersion(version string, releaseType string) string {
	parts := strings.SplitN(version, ".", 3)

	major, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	minor, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	patch, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	if releaseType == "major" {
		major = major + 1
		minor = 0
		patch = 0
	} else if releaseType == "minor" {
		minor = minor + 1
		patch = 0
	} else {
		patch = patch + 1
	}

	newVersion := fmt.Sprintf("%d.%d.%d", major, minor, patch)
	return newVersion
}
