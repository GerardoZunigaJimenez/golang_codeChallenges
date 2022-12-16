package main

import (
	"globant_golang_codeChallenges/deuna/endpoint"
	"net/http"
)

func main() {
	http.HandleFunc("/api/v1/starships/available?passengers=500", endpoint.StarshipsAvailable)
	http.ListenAndServe(":3000", nil)
}
