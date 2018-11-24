package release

import (
	"bytes"
	"deployer/pkg"
	"fmt"
	"html/template"
	"os"
	"strconv"

	"github.com/jmcvetta/napping"
)

func isGitlabTokenAdmin() bool {
	gitlabTokenAdmin := os.Getenv("DEPLOYER_IS_GITLAB_TOKEN_ADMIN")
	if gitlabTokenAdmin == "" {
		pkg.FatalF("%s\n", "Environment variable DEPLOYER_IS_GITLAB_TOKEN_ADMIN not set")
	}
	b, err := strconv.ParseBool(gitlabTokenAdmin)
	if err != nil {
		pkg.FatalF("%s", err)
	}
	return b
}

func createTag(projectID string, userID string, tag string, ref string, changeLog string) {
	endPoint := fmt.Sprintf("/projects/%s/repository/tags", projectID)

	headers := make(map[string]string)
	if isGitlabTokenAdmin() {
		headers["SUDO"] = userID
	}

	payload := napping.Params{
		"tag_name":            tag,
		"ref":                 ref,
		"release_description": changeLog,
	}.AsUrlValues().Encode()

	endPoint = endPoint + "?" + payload

	apiResponse := make(map[string]interface{})
	err := pkg.CallGitlabAPI(endPoint, headers, nil, &apiResponse)
	if err != nil {
		pkg.FatalF("Could not create tag %v\n", err.Error())
	}

	message, ok := apiResponse["message"]
	if !ok || message == nil {
		pkg.FatalF("Could not create tag, error message: %v\n", message)
	}
}

func gitlabReleaseNotes(changes []string, releaserName string) string {
	tmpl, err := template.New("release").Parse(pkg.BodyTemplate)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
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
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
	return out.String()
}

func mergeRequestDescription(promoter string) string {
	description := ""
	if !isGitlabTokenAdmin() && promoter != "" {
		description = fmt.Sprintf("Promoted by **%s**", promoter)
	}

	return description
}

func createMR(sourceBranch string, targetBranch string, projectID string, userID string, promoter string) {
	endPoint := fmt.Sprintf("/projects/%s/merge_requests", projectID)

	headers := make(map[string]string)
	if isGitlabTokenAdmin() {
		headers["SUDO"] = userID
	}

	payload := napping.Params{
		"source_branch":        sourceBranch,
		"title":                fmt.Sprintf("Promoting %s to %s", sourceBranch, targetBranch),
		"target_branch":        targetBranch,
		"remove_source_branch": "false",
		"labels":               "promote, pipeline",
		"description":          mergeRequestDescription(promoter),
	}.AsUrlValues().Encode()

	endPoint = endPoint + "?" + payload

	apiResponse := make(map[string]interface{})
	err := pkg.CallGitlabAPI(endPoint, headers, nil, &apiResponse)
	if err != nil {
		pkg.FatalF("Could not create tag %v\n", err.Error())
	}

	message, ok := apiResponse["message"]
	if ok && message != nil {
		pkg.FatalF("Could not create tag, error message: %v\n", message)
	}
}
