package s3

import (
	"deployer/pkg"
	"fmt"
	"os"
	"strings"
)

func ExportDockerImageToS3(repo string, tag string) (*UploadDetails, error) {
	repoWithTag := fmt.Sprintf("%s:%s", repo, tag)

	r := strings.Split(repo, "/")
	projectName := r[len(r)-1]
	outputFilePath := fmt.Sprintf("/tmp/%s-%s.tar", projectName, tag)

	command := fmt.Sprintf("docker image pull %s", repoWithTag)
	fmt.Println(command, " \n")
	err := pkg.Execute(command)
	if err != nil {
		return nil, err
	}

	command = fmt.Sprintf("docker image save -o %s %s", outputFilePath, repoWithTag)
	fmt.Println(command, " \n")
	err = pkg.Execute(command)
	if err != nil {
		return nil, err
	}
	defer os.RemoveAll(outputFilePath)

	err = Init()
	if err != nil {
		return nil, err
	}
	return Upload(outputFilePath, true)
}
