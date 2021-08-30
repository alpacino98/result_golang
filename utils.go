package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func makeRequests(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(router)
	goroutine, err := strconv.Atoi(params["goroutine"])
	if err != nil {
		resp := make(map[string]string)
		resp["message"] = "Something went wrong"
		jsonResp, err := json.Marshal(resp)
		if err != nil {
			log.Fatalf("Error happened in JSON marshal. Err: %s", err)
		}
		w.Write(jsonResp)
		return
	}
	makeResultRequests(goroutine)
	json.NewEncoder(w).Encode(lastResults)
}

func getMain(w http.ResponseWriter, router *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	resp := make(map[string]string)
	resp["message"] = "Make request to /api/request or /api/results"
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		log.Fatalf("Error happened in JSON marshal. Err: %s", err)
	}
	w.Write(jsonResp)
	return
}

func getResults(w http.ResponseWriter, router *http.Request) {
	log.Print(results)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}
