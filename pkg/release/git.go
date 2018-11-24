package release

import (
	"errors"
	"os/exec"
	"strings"
)

func lastTag() (string, error) {
	tag, err := cleanGit("describe", "--tags", "--abbrev=0")
	return tag, err
}

func gitLog(prevTag string) (string, error) {
	refs := "HEAD"
	if prevTag != "" {
		refs = prevTag + ".." + refs
	}
	var args = []string{"log", "--pretty=oneline", "--abbrev-commit"}
	args = append(args, refs)
	return git(args...)
}

func gitSHA(tag string) (string, error) {
	sha, err := cleanGit("rev-parse", tag)
	return sha, err
}

func cleanGit(args ...string) (output string, err error) {
	output, err = git(args...)
	return strings.Replace(strings.Split(output, "\n")[0], "'", "", -1), err
}

func git(args ...string) (output string, err error) {
	var cmd = exec.Command("git", args...)
	bts, err := cmd.CombinedOutput()
	if err != nil {
		return "", errors.New(string(bts))
	}
	return string(bts), err
}
