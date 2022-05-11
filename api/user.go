package api

import (
	"io/ioutil"
	"log"
	"net/http"
)

func GetRandomUser(w http.ResponseWriter, r *http.Request) {
	apiURL := "https://randomuser.me/api"

	req, err := http.NewRequest(http.MethodGet, apiURL, nil)
	if err != nil {
		log.Println(err)
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)

	w.WriteHeader(http.StatusOK)
	w.Write(data)
}
