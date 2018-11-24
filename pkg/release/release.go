package release

import (
	"deployer/pkg"
	"fmt"
	"os"
	"strings"
)

func CheckForGitRepo() {
	_, err := os.Stat(".git")
	if err != nil {
		fmt.Println("Git repo not found")
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
	fmt.Println("Git repo found...")
}

func Create(releaseType string, projectID string, userID string, releaserName string) {
	latestRelease := "0.0.0"
	tag, err := lastTag()
	if err != nil {
		tag = ""
		fmt.Println("No prior releases found")
	} else {
		// removes usual `v` prefix
		latestRelease = strings.TrimPrefix(tag, "v")
		fmt.Println("Last release is " + latestRelease)
	}

	gitLog, err := gitLog(tag)
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}

	gitLog = strings.TrimSpace(gitLog)
	gitLog = strings.Replace(gitLog, "'", "", -1)
	changes := strings.Split(gitLog, "\n")

	numChanges := len(changes)
	if numChanges <= 1 {
		pkg.FatalF("No changes to the last release detected \n")
	}
	if numChanges > pkg.MaxChanges {
		changes = changes[:pkg.MaxChanges]
		changes = append(changes, "Truncated")
	}

	description := gitlabReleaseNotes(changes, releaserName)

	newVersion := newVersion(latestRelease, releaseType)
	fmt.Println("Creating release", newVersion+"...")
	newTag := "v" + newVersion
	ref, err := gitSHA("HEAD")
	if err != nil {
		pkg.FatalF("An error occurred:\n %s \n", err.Error())
	}
	createTag(projectID, userID, newTag, ref, description)
}

func Promote(sourceBranch string, targetBranch string, projectID string, user string, promoter string) {
	createMR(sourceBranch, targetBranch, projectID, user, promoter)
}
