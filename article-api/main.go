package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func homePage(w http.ResponseWriter, r *http.Request){
    fmt.Fprintf(w, "Welcome to the HomePage!")
    fmt.Println("Endpoint Hit: homePage")
}

type Article struct {
    Title string `json:"Title"`
    Desc string `json:"desc"`
    Content string `json:"content"`
}

// let's declare a global Articles array
// that we can then populate in our main function
// to simulate a database
var Articles []Article

func handleRequests() {
    http.HandleFunc("/", homePage)
    // add our articles route and map it to our 
    // returnAllArticles function like so
    http.HandleFunc("/articles", returnAllArticles)
    log.Fatal(http.ListenAndServe(":10000", nil))
}

func returnAllArticles(w http.ResponseWriter, r *http.Request){
    fmt.Println("Endpoint Hit: returnAllArticles")
    json.NewEncoder(w).Encode(Articles)
}

func main() {
	Articles = []Article{
        {Title: "Hello", Desc: "Article Description", Content: "Article Content"},
        {Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
    }
    handleRequests()
}