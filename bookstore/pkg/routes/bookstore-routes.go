package routes

import (
	"github.com/gorilla/mux"
	"github.com/saikumar-3093/go-projects/bookstore/pkg/controllers"
)

func RegisterBookStoreRoute(router *mux.Router) {

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookbyID).Methods("GET")
	// router.HandleFunc("/book/{bookId}", controllers.UpdateBookbyID).Methods("PUT")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBookbyID).Methods("DELETE")
}
