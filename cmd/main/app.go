package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"strconv"
)

type Query struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbin"`
	Prompt string `json:"prompt"`
	Seed   string `json:"seed"`
}

var queries []Query

func getQueries(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(queries)
	if err != nil {
		return
	}

}

func getQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for _, item := range queries {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)

			return
		}
	}

}

func deleteQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range queries {
		if item.ID == params["id"] {
			queries = append(queries[:index], queries[index+1:]...)
			break
		}
	}
	err := json.NewEncoder(w).Encode(queries)
	if err != nil {
		return
	}
}

func createQuery(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var query Query
	err := json.NewDecoder(r.Body).Decode(&query)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	query.ID = strconv.Itoa(rand.Intn(10000000))
	queries = append(queries, query)
	err = json.NewEncoder(w).Encode(query)
	if err != nil {
		return
	}

}

func main() {
	r := mux.NewRouter()
	queries = append(queries, Query{ID: "1", Isbn: "438227", Prompt: "White shirt", Seed: "39"})
	queries = append(queries, Query{ID: "2", Isbn: "438227", Prompt: "Blue elephant", Seed: "239"})
	r.HandleFunc("/queries", getQueries).Methods("GET")
	r.HandleFunc("/queries/{id}", getQuery).Methods("GET")
	r.HandleFunc("/queries", createQuery).Methods("POST")
	r.HandleFunc("/queries/{id}", deleteQuery).Methods("DELETE")
	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8800", r))
}
