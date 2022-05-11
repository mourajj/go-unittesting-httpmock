package main

import (
	"log"
	"net/http"
	"unittest/api"
)

func main() {

	http.HandleFunc("/", api.GetRandomUser)
	log.Println("listening...")
	panic(http.ListenAndServe("localhost:7000", nil))
}
