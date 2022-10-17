package routes

import (
	"github.com/gorilla/mux" 
	controller "books/pkg/controllers"
)

var BookRoutes = func(router *mux.Router) {
	router.HandleFunc("/book/" , controller.Index).Methods("GET")
	router.HandleFunc("/book/" , controller.Create).Methods("POST")
	router.HandleFunc("/book/{bookId}" , controller.Show).Methods("GET")
	router.HandleFunc("/book/{bookId}" , controller.Edit).Methods("PUT")
	router.HandleFunc("/book/{bookId}" , controller.Delete).Methods("DELETE")
}