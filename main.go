package main

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
)

type ApiResponse struct {
	Message string `json:"message"`
}

func main() {
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/", homePage)
	http.HandleFunc("/login", loginPage)
	log.Fatal(http.ListenAndServe(":8081", nil))
}

func homePage(res http.ResponseWriter, req *http.Request) {
	apiResponse := ApiResponse {
		Message:"Welcome to home page",
	}
	res.Header().Set("Content-Type", "text/json")
	fmt.Println("Endpoint hit")
	json.NewEncoder(res).Encode(apiResponse)
}

func loginPage(res http.ResponseWriter, req *http.Request) {
	apiResponse := ApiResponse {
		Message:"Welcome to login page",
	}
	res.Header().Set("Content-Type", "text/json")
	fmt.Println("Endpoint hit")
	json.NewEncoder(res).Encode(apiResponse)
}