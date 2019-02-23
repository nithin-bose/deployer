package release

import (
	"fmt"
	"strconv"
	"strings"
)

func newVersion(version string, releaseType string) (string, error) {
	parts := strings.SplitN(version, ".", 3)

	major, err := strconv.ParseUint(parts[0], 10, 64)
	if err != nil {
		return "", err
	}

	minor, err := strconv.ParseUint(parts[1], 10, 64)
	if err != nil {
		return "", err
	}

	patch, err := strconv.ParseUint(parts[2], 10, 64)
	if err != nil {
		return "", err
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
	return newVersion, nil
}
