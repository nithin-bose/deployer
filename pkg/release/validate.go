package release

import "errors"

func ValidateType(release string) error {
	if release != "major" && release != "minor" && release != "patch" {
		return errors.New("Invalid release type. Only major, minor and patch are supported")
	}
	return nil
}
