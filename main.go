package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	ListRouters(router)

	fmt.Println("Server will start at http://localhost:8000/")
	log.Fatal(http.ListenAndServe(":8000", router))
}
