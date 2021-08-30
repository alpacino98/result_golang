package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

type Results struct {
	Websites       []websites   `json:"titles"`
	Total_requests total_result `json:"total_result"`
}

type total_result struct {
	Success   int `json:"succesful_total"`
	Unsuccess int `json:"unsuccesful_total"`
}

type websites struct {
	Url       string `json:"url"`
	Title     string `json:"title"`
	Success   int    `json:"succesful_requests"`
	Unsuccess int    `json:"unsuccesful_requests"`
}

var results Results
var lastResults map[string]bool = make(map[string]bool)

func main() {
	//init jsonBody, and dataholder
	results.Websites = append(results.Websites, websites{Url: "https://www.result.si/projekti/", Title: "", Success: 0, Unsuccess: 0})
	results.Websites = append(results.Websites, websites{Url: "https://www.result.si/o-nas/", Title: "", Success: 0, Unsuccess: 0})
	results.Websites = append(results.Websites, websites{Url: "https://www.result.si/kariera/", Title: "", Success: 0, Unsuccess: 0})
	results.Websites = append(results.Websites, websites{Url: "https://www.result.si/blog/", Title: "", Success: 0, Unsuccess: 0})
	results.Total_requests.Success = 0
	results.Total_requests.Unsuccess = 0
	// Init router
	router := mux.NewRouter()

	//Router handler / Endpoints
	router.HandleFunc("/api/", getMain).Methods("GET")
	router.HandleFunc("/api/request/{goroutine}", makeRequests).Methods("GET")
	router.HandleFunc("/api/results", getResults).Methods("GET")
	fmt.Println("server started")
	http.ListenAndServe(":8080", router)
}
