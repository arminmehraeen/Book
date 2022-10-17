package main

import (
	"books/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	routes.BookRoutes(r)
	http.Handle("/" , r)
	fmt.Println("Server run . ")
	log.Fatal(http.ListenAndServe("localhost:5353" , r))
}