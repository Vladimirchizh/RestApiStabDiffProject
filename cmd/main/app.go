package main

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"
)

type Query struct {
	ID     string `json:"id"`
	Isbn   string `json:"isbin"`
	Prompt string `json:"prompt"`
	Seed   string `json:"seed"`
}

var (
	queries       []Query
	WarningLogger *log.Logger
	InfoLogger    *log.Logger
	ErrorLogger   *log.Logger
)

func init() {
	file, err := os.OpenFile("logs.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatal(err)
	}

	InfoLogger = log.New(file, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	WarningLogger = log.New(file, "WARNING: ", log.Ldate|log.Ltime|log.Lshortfile)
	ErrorLogger = log.New(file, "ERROR: ", log.Ldate|log.Ltime|log.Lshortfile)
}

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
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				return
			}
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
			WarningLogger.Println("Item was deleted, ID: ", item.ID)
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
	InfoLogger.Println("New item inserted: ", query)
}

func main() {
	r := mux.NewRouter()
	startingMessage := "Starting server at port 8000\n"

	queries = append(queries, Query{ID: "1", Isbn: "438227", Prompt: "White shirt", Seed: "39"})
	queries = append(queries, Query{ID: "2", Isbn: "438227", Prompt: "Blue elephant", Seed: "239"})
	r.HandleFunc("/queries", getQueries).Methods("GET")
	r.HandleFunc("/queries/{id}", getQuery).Methods("GET")
	r.HandleFunc("/queries", createQuery).Methods("POST")
	r.HandleFunc("/queries/{id}", deleteQuery).Methods("DELETE")
	fmt.Printf(startingMessage)
	InfoLogger.Println(startingMessage)
	ErrorLogger.Fatalln(http.ListenAndServe(":8800", r))
	// log.Fatal(http.ListenAndServe(":8800", r))
}
