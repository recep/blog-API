package routes

import (
	"github.com/gorilla/mux"

	. "blog-API/handlers"
)

func ReturnRoutes(router *mux.Router) {
	router.HandleFunc("/blogs", ReturnAllPost).Methods("GET")
	router.HandleFunc("/blogs", CreateNewPost).Methods("POST")
	router.HandleFunc("/blogs/{id}", ReturnSinglePost).Methods("GET")
	router.HandleFunc("/blogs/{id}", UpdatePost).Methods("PUT")
	router.HandleFunc("/blogs/{id}", DeletePost).Methods("DELETE")
}
