package webhook

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func Run(socket bool) {
	r := mux.NewRouter()
	// Routes consist of a path and a handler function.
	r.HandleFunc("/", HomeHandler).Methods("GET")
	r.HandleFunc("/docker/deploy/app", DockerDeployAppHandler).Methods("POST")
	r.HandleFunc("/k8s/deploy/app", K8sDeployAppHandler).Methods("POST")

	if !socket {
		port := os.Getenv("DEPLOYER_WEBHOOK_PORT")
		if port == "" {
			port = "3000"
		}
		fmt.Println("Listening on port", port+"...")
		log.Fatal(http.ListenAndServe(":"+port, r))
	} else {
		sock := "/var/run/deployer.sock"
		os.Remove(sock)
		unixListener, err := net.Listen("unix", sock)
		if err != nil {
			log.Fatal("Error listening on socket: ", err)
		}
		defer unixListener.Close()
		fmt.Println("Listening on socket", sock+"...")
		log.Fatal(http.Serve(unixListener, r))
	}
}
