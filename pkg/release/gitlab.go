package release

import (
	"bytes"
	"deployer/pkg"
	"errors"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/jmcvetta/napping"
)

func isGitlabTokenAdmin() (bool, error) {
	gitlabTokenAdmin := os.Getenv("DEPLOYER_IS_GITLAB_TOKEN_ADMIN")
	if gitlabTokenAdmin == "" {
		return false, errors.New("Environment variable DEPLOYER_IS_GITLAB_TOKEN_ADMIN not set")
	}
	b, err := strconv.ParseBool(gitlabTokenAdmin)
	if err != nil {
		return false, err
	}
	return b, nil
}

func createTag(projectID string, userID string, tag string, ref string, changeLog string) error {
	endPoint := fmt.Sprintf("/projects/%s/repository/tags", projectID)

	headers := make(map[string]string)
	gitlabAdmin, err := isGitlabTokenAdmin()
	if err != nil {
		return err
	}
	if gitlabAdmin {
		headers["SUDO"] = userID
	}

	payload := napping.Params{
		"tag_name":            tag,
		"ref":                 ref,
		"release_description": changeLog,
	}.AsUrlValues().Encode()

	endPoint = endPoint + "?" + payload

	apiResponse := make(map[string]interface{})
	err = pkg.CallGitlabAPI(endPoint, headers, nil, &apiResponse)
	if err != nil {
		return err
	}

	message, ok := apiResponse["message"]
	if !ok || message == nil {
		return errors.New("Could not create tag, " + message.(string))
	}
	return nil
}

func gitlabReleaseNotes(changes []string, releaserName string) (string, error) {
	tmpl, err := template.New("release").Parse(pkg.BodyTemplate)
	if err != nil {
		return "", err
	}

	var out bytes.Buffer
	err = tmpl.Execute(&out, struct {
		Changes      []string
		ReleaserName string
	}{
		Changes:      changes,
		ReleaserName: releaserName,
	})

	if err != nil {
		return "", err
	}
	return out.String(), nil
}

func mergeRequestDescription(promoter string) (string, error) {
	description := ""
	gitlabAdmin, err := isGitlabTokenAdmin()
	if err != nil {
		return "", err
	}
	if !gitlabAdmin && promoter != "" {
		description = fmt.Sprintf("Promoted by **%s**", promoter)
	}

	return description, nil
}

func createMR(sourceBranch string, targetBranch string, projectID string, userID string, promoter string) error {
	endPoint := fmt.Sprintf("/projects/%s/merge_requests", projectID)

	headers := make(map[string]string)
	gitlabAdmin, err := isGitlabTokenAdmin()
	if err != nil {
		return err
	}
	if gitlabAdmin {
		headers["SUDO"] = userID
	}

	description, err := mergeRequestDescription(promoter)
	if err != nil {
		return err
	}

	payload := napping.Params{
		"source_branch":        sourceBranch,
		"title":                fmt.Sprintf("Promoting %s to %s", sourceBranch, targetBranch),
		"target_branch":        targetBranch,
		"remove_source_branch": "false",
		"labels":               "promote, pipeline",
		"description":          description,
	}.AsUrlValues().Encode()

	endPoint = endPoint + "?" + payload

	apiResponse := make(map[string]interface{})
	err = pkg.CallGitlabAPI(endPoint, headers, nil, &apiResponse)
	if err != nil {
		return err
	}

	message, ok := apiResponse["message"]
	if ok && message != nil {
		return errors.New("Could not create tag, " + message.(string))
	}
	return nil
}
