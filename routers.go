package main

import (
	"net/http"

	"github.com/gorilla/mux"
)

// ListRouters : adalah function untuk menghandle router pada apps
func ListRouters(router *mux.Router) {

	// Users Route
	router.HandleFunc("/users", returnAllUsers).Methods("GET")
	router.HandleFunc("/users", insertUsersMultipart).Methods("POST")
	router.HandleFunc("/users", updateUsersMultipart).Methods("PUT")
	router.HandleFunc("/users", deleteUsersMultipart).Methods("DELETE")

	http.Handle("/", router)
}
