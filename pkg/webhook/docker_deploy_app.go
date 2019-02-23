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
	composeFile := os.Getenv("DEPLOYER_DEFAULT_COMPOSE_FILE")
	if composeFile == "" {
		err = errors.New("DEPLOYER_DEFAULT_COMPOSE_FILE not set")
		RenderError(w, err)
		return
	}

	composeFileDir := os.Getenv("DEPLOYER_DEFAULT_COMPOSE_FILE_DIR")
	if composeFileDir == "" {
		err = errors.New("DEPLOYER_DEFAULT_COMPOSE_FILE_DIR not set")
		RenderError(w, err)
		return
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

	err = os.Chdir(composeFileDir)
	if err != nil {
		RenderError(w, err)
		return
	}
	err = docker.DeployServiceApp(composeFile, req.App)
	if err != nil {
		RenderError(w, err)
		return
	}
	RenderSuccess(w, nil)
}
