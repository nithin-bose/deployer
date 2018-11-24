package webhook

import "net/http"

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	data := make(map[string]interface{})
	data["message"] = "Up and running"
	RenderSuccess(w, data)
}
