package deploy

import "errors"

func ValidateEnvironment(environment string) error {
	if environment != "staging" && environment != "production" {
		return errors.New("Unknown environment. Supported environments are staging and production")
	}
	return nil
}
