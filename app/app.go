package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	router := mux.NewRouter().StrictSlash(true)

	router.Handle("/", http.HandlerFunc(serve)).Methods("GET")
	router.Handle("/metrics", promhttp.Handler())

	log.Fatal(http.ListenAndServe(":80", router))
}

type response struct {
	Status string
	Code   int
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response{Status: "OK", Code: http.StatusOK})

	w.WriteHeader(http.StatusOK)
}
