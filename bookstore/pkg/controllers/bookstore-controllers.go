package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/saikumar-3093/go-projects/bookstore/pkg/models"
	"github.com/saikumar-3093/go-projects/bookstore/pkg/utils"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	newBook := models.GetAllBooks()

	resp, err := json.Marshal(newBook)
	if err != nil {
		log.Fatal("error while Marshaling data")
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func GetBookbyID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["bookId"], 10, 64)
	if err != nil {
		log.Fatal("error while converting string to int")
	}
	book, _ := models.GetBookbyID(id)
	resp, err := json.Marshal(&book)
	if err != nil {
		log.Fatal("error while Marshaling data")
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	newbook := &models.Book{}
	utils.ParseBody(r, newbook)
	book := newbook.CreateBook()

	resp, err := json.Marshal(book)
	if err != nil {
		log.Fatal("error while Marshaling data")
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)
}

func DeleteBookbyID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.ParseInt(vars["bookId"], 10, 64)
	if err != nil {
		log.Fatal("error while converting string to int")
	}
	book := models.DeleteBookbyID(id)
	resp, err := json.Marshal(book)
	if err != nil {
		log.Fatal("error while Marshaling data")
	}
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(resp)

}

// func UpdateBookbyID(w http.ResponseWriter, r *http.Request) {
// 	vars := mux.Vars(r)
// 	var updateBook = &models.Book{}
// 	utils.ParseBody(r, updateBook)
// 	id, err := strconv.ParseInt(vars["bookId"], 10, 64)
// 	if err != nil {
// 		log.Fatal("error while converting string to int")
// 	}
// 	book, db := models.UpdateBookbyID(id, updateBook)
// 	db.Save(&book)
// 	resp, err := json.Marshal(book)
// 	if err != nil {
// 		log.Fatal("error while Marshaling data")
// 	}
// 	w.Header().Set("Content-Type", "application/json")
// 	w.WriteHeader(http.StatusOK)
// 	w.Write(resp)
// }
