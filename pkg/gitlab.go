package pkg

import (
	"errors"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/jmcvetta/napping"
)

func getGitlabToken() string {
	gitlabToken := os.Getenv("DEPLOYER_GITLAB_ACCESS_TOKEN")
	if gitlabToken == "" {
		log.Fatal("Gitlab token not set")
	}
	return gitlabToken
}

func callAPI(server string, endPoint string, headers map[string]string, payload interface{}, apiResponse interface{}) error {
	s := napping.Session{}
	url := server + endPoint

	if headers != nil {
		s.Header = &http.Header{}
		for k, v := range headers {
			s.Header.Set(k, v)
		}
	}

	resp, err := s.Post(url, payload, apiResponse, nil)
	if err != nil {
		return err
	}

	if resp.Status() != 200 && resp.Status() != 201 {
		return errors.New("Got Status " + strconv.Itoa(resp.Status()) + ". Expected 200 or 201 OK")
	}

	return nil
}

func CallGitlabAPI(endPoint string, headers map[string]string, payload interface{}, apiResponse interface{}) error {
	headers["PRIVATE-TOKEN"] = getGitlabToken()
	return callAPI(GitlabServer, endPoint, headers, payload, apiResponse)
}
