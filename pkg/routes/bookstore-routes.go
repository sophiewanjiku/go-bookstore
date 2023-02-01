package main

import (
	"github.com/gorilla/mux"
	"github.com/sophiewanjiku/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("UPDATE")
	router.HandleFunc("/book{bookId}", controllers.DeleteBook).Methods("DELETE")
}
