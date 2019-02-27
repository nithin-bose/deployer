package webhook

import (
	"deployer/pkg/deploy/k8s"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
)

func K8sDeployAppHandler(w http.ResponseWriter, r *http.Request) {
	chartsDir := os.Getenv("DEPLOYER_HELM_CHARTS_DIR")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		RenderError(w, err)
		return
	}
	var req K8sRequest
	err = json.Unmarshal(body, &req)
	if err != nil {
		RenderError(w, err)
		return
	}
	err = authenticateK8s(req.AccessKey, req.AccessToken)
	if err != nil {
		RenderError(w, err)
		return
	}

	if req.Environment == "" {
		err = errors.New("Required field 'environment' cannot be empty")
		RenderError(w, err)
		return
	}

	if req.App == "" {
		err = errors.New("Required field 'app' cannot be empty")
		RenderError(w, err)
		return
	}

	if req.Version == "" {
		err = errors.New("Required field 'version' cannot be empty")
		RenderError(w, err)
		return
	}

	err = k8s.DeployServiceApp(chartsDir, false, true, req.Environment, req.App, req.Version)
	if err != nil {
		RenderError(w, err)
		return
	}
	RenderSuccess(w, nil)
}
