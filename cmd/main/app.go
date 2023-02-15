package main

import (
	"RestApiStabDiffProject/internal/queries"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func main() {
	router := mux.NewRouter()
	fmt.Printf("Starting server at port 8800\n")
	handler := queries.NewHandler()
	handler.Register(router)

	log.Fatal(http.ListenAndServe(":8800", router))
}
