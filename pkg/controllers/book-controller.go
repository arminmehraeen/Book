package controllers

import (
	"books/pkg/models"
	"books/pkg/utils"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func Index(w http.ResponseWriter , r *http.Request) {
	newBooks := models.Index()
	res , _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type" , "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Show(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id , err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing ( show method )")
	}

	book , _ := models.Show(Id)

	res , _ := json.Marshal(book)
	w.Header().Set("Content-Type" , "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Create(w http.ResponseWriter , r *http.Request) {
	book := &models.Book{}
	utils.ParseBody(r,book)
	b := book.Create()
	
	res , _ := json.Marshal(b)
	w.Header().Set("Content-Type" , "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


func Delete(w http.ResponseWriter , r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id , err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing ( show method )")
	}

	book := models.Delete(Id)

	res , _ := json.Marshal(book)
	w.Header().Set("Content-Type" , "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func Edit(w http.ResponseWriter , r *http.Request) {
	var book = &models.Book{}
	utils.ParseBody(r,book)
	vars := mux.Vars(r)
	bookId := vars["bookId"]

	Id , err := strconv.ParseInt(bookId,0,0)
	if err != nil {
		fmt.Println("error while parsing ( show method )")
	}

	bookDetails , db := models.Show(Id)

	if book.Name != "" {
		bookDetails.Name = book.Name
	}

	if book.Author != "" {
		bookDetails.Author = book.Author
	}

	if book.Publication != "" {
		bookDetails.Publication = book.Publication
	}

	db.Save(&bookDetails)

	res , _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type" , "aplication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}


