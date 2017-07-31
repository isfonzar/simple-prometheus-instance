package main

import (
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"encoding/json"
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
	Code int
}

func serve(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	json.NewEncoder(w).Encode(response{Status: "OK", Code: 200})

	w.WriteHeader(http.StatusOK)
}