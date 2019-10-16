package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
)

type article struct {
	ID      string `json:"ID"`
	Title   string `json:"Title"`
	Content string `json:"Content"`
}

var articles []article = []article{{"10", "My story", "This is all the content i can think of"}, {"20", "Your story", "This is all the content i can think of"}}

func helloFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello world")
}

func allArticles(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(articles)
}

func myArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	found := false
	for _, article := range articles {
		if article.ID == id {
			json.NewEncoder(w).Encode(article)
			found = true
		}
	}
	if !found {
		fmt.Fprint(w, "Article not found")
	}
}

func createArticle(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newArticle article
	json.Unmarshal(reqBody, &newArticle)
	articles = append(articles, newArticle)
	json.NewEncoder(w).Encode(articles)
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	found := false
	for i, article := range articles {
		if article.ID == id {
			articles = append(articles[:i], articles[i+1:]...)
			fmt.Fprint(w, "Article deleted")
			found = true
		}
	}
	if !found {
		fmt.Fprint(w, "Article not present")
	}
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	found := false
	reqBody, _ := ioutil.ReadAll(r.Body)
	var newArticle article
	json.Unmarshal(reqBody, &newArticle)
	for i, article := range articles {
		if article.ID == id {
			newArticle.ID = id
			articles[i] = newArticle
			json.NewEncoder(w).Encode(articles[i])
			found = true
		}
	}
	if !found {
		fmt.Fprint(w, "Article not present")
	}
}

func main() {

	mymux := mux.NewRouter().StrictSlash(true)
	mymux.HandleFunc("/", helloFunc)
	mymux.HandleFunc("/articles", allArticles)
	mymux.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	mymux.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	mymux.HandleFunc("/article/{id}", myArticle)
	mymux.HandleFunc("/article", createArticle).Methods("POST")

	http.ListenAndServe(":9090", mymux)
}
