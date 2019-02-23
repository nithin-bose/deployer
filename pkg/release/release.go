package release

import (
	"deployer/pkg"
	"errors"
	"fmt"
	"os"
	"strings"
)

func CheckForGitRepo() bool {
	_, err := os.Stat(".git")
	if err != nil {
		return false
	}
	return true
}

func Create(releaseType string, projectID string, userID string, releaserName string) error {
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
		return err
	}

	gitLog = strings.TrimSpace(gitLog)
	gitLog = strings.Replace(gitLog, "'", "", -1)
	changes := strings.Split(gitLog, "\n")

	numChanges := len(changes)
	if numChanges <= 1 {
		return errors.New("No changes to the last release detected")
	}
	if numChanges > pkg.MaxChanges {
		changes = changes[:pkg.MaxChanges]
		changes = append(changes, "Truncated")
	}

	description, err := gitlabReleaseNotes(changes, releaserName)
	if err != nil {
		return err
	}

	newVersion, err := newVersion(latestRelease, releaseType)
	if err != nil {
		return err
	}
	fmt.Println("Creating release", newVersion+"...")
	newTag := "v" + newVersion
	ref, err := gitSHA("HEAD")
	if err != nil {
		return err
	}
	return createTag(projectID, userID, newTag, ref, description)
}

func Promote(sourceBranch string, targetBranch string, projectID string, user string, promoter string) error {
	return createMR(sourceBranch, targetBranch, projectID, user, promoter)
}
