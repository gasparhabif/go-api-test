package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strings"
)

const PORT = ":10000"

type Article struct {
	Id				string `json:"id"`
	Title 			string `json:"Title"`
	Description 	string `json:"description"`
	Content 		string `json:"content"`
}

var Articles []Article

func main() {
	fmt.Println("Starting API")
	Articles = []Article{
		Article{Id: "1", Title: "Hello", Description: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Description: "Article Description", Content: "Article Content"},
	}
	handleRequests()
}

func handleRequests() {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", homePage)
	router.HandleFunc("/articles", allArticles)
	router.HandleFunc("/article/{id}", singleArticle)

	log.Fatal(http.ListenAndServe(PORT, router))
}

func printCurrentEndpoint(endpoint string) {
	if !strings.Contains(endpoint, "/favicon.ico") {
		fmt.Println("Endpoint: " + endpoint)
	}
}

func homePage(w http.ResponseWriter, r *http.Request){
	printCurrentEndpoint(r.URL.String())
	fmt.Fprintf(w, "Holis")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	printCurrentEndpoint(r.URL.String())
	json.NewEncoder(w).Encode(Articles)
}

func singleArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	printCurrentEndpoint(r.URL.String())

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}