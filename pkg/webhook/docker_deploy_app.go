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
	err = authenticate(req.AccessKey, req.AccessToken)
	if err != nil {
		RenderError(w, err)
		return
	}

	if req.App == "" {
		err = errors.New("Required field 'app' cannot be empty")
		RenderError(w, err)
		return
	}

	if req.ComposeFile == "" {
		req.ComposeFile = os.Getenv("DEPLOYER_DEFAULT_COMPOSE_FILE")
		if req.ComposeFile == "" {
			err = errors.New("Required field 'compose_file' cannot be empty")
			RenderError(w, err)
			return
		}
	}
	docker.DeployServiceApp(req.ComposeFile, req.App)
	if r := recover(); r != nil {
		err := errors.New("Deployment failed")
		RenderError(w, err)
		return
	}
	RenderSuccess(w, nil)
}
