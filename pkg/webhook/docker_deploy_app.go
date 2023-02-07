package webhook

import (
	"deployer/pkg/deploy/docker"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

func DockerDeployAppHandler(w http.ResponseWriter, r *http.Request) {
	var err error
	composeFile := os.Getenv("DEPLOYER_COMPOSE_FILE")
	if composeFile == "" {
		composeFile = "docker-compose.yml"
	}

	dockerStacksDir := os.Getenv("DEPLOYER_DOCKER_STACKS_DIR")
	if dockerStacksDir == "" {
		dockerStacksDir = "/root/docker-stacks"
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RenderError(w, err)
		return
	}
	var req DockerRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		RenderError(w, err)
		return
	}
	err = authenticateDocker(req.AccessKey, req.AccessToken)
	if err != nil {
		RenderError(w, err)
		return
	}

	if req.App == "" {
		err = errors.New("Required field 'app' cannot be empty")
		RenderError(w, err)
		return
	}

	composeFileDir := dockerStacksDir + string(os.PathSeparator) + req.App
	err = os.Chdir(composeFileDir)
	if err != nil {
		RenderError(w, err)
		return
	}

	err = docker.DeployServiceApp(composeFile, req.Service)
	if err != nil {
		RenderError(w, err)
		return
	}
	RenderSuccess(w, nil)
}
