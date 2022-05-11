package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

type HTTPClientInterdace interface {
	Do(req *http.Request) (*http.Response, error)
}

type Service struct {
	HTTPClient HTTPClientInterdace
}

func (s *Service) GetRandomUser(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://randomuser.me/api"

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		log.Println(err)
	}

	//res, err := http.DefaultClient.Do(req)
	res, err := s.HTTPClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
	}

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
