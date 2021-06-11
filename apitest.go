package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io/ioutil"
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
	router.HandleFunc("/articles", allArticles).Methods("GET")
	router.HandleFunc("/article/{id}", singleArticle).Methods("GET")
	router.HandleFunc("/article", addArticle).Methods("POST")
	router.HandleFunc("/article/{id}", removeArticle).Methods("DELETE")
	router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")


	log.Fatal(http.ListenAndServe(PORT, router))
}

func printCurrentEndpoint(endpoint string) {
	if !strings.Contains(endpoint, "/favicon.ico") {
		fmt.Println("Endpoint: " + endpoint)
	}
}

func homePage(w http.ResponseWriter, request *http.Request){
	printCurrentEndpoint(request.URL.String())
	fmt.Fprintf(w, "Holis")
}

func allArticles(w http.ResponseWriter, request *http.Request) {
	printCurrentEndpoint(request.URL.String())
	json.NewEncoder(w).Encode(Articles)
}

func singleArticle(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	key := vars["id"]

	printCurrentEndpoint(request.URL.String())

	for _, article := range Articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article)
		}
	}
}

func addArticle(w http.ResponseWriter, request *http.Request) {
	reqBody, _ := ioutil.ReadAll(request.Body)
	var article Article
	json.Unmarshal(reqBody, &article)
	Articles = append(Articles, article)

	json.NewEncoder(w).Encode(article)
}

func removeArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range Articles {
		if article.Id == id {
			Articles = append(Articles[:index], Articles[index+1:]...)
		}
	}
}

func updateArticle(w http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(request.Body)
	var updatedArticle Article
	json.Unmarshal(reqBody, &updatedArticle)

	for index, article := range Articles {
		if article.Id == id {
			var articleToUpdate = Articles[index]
			articleToUpdate.Title = updatedArticle.Title
			articleToUpdate.Content = updatedArticle.Content
			articleToUpdate.Description = updatedArticle.Description
			Articles[index] = articleToUpdate
			json.NewEncoder(w).Encode(articleToUpdate)
		}
	}
}