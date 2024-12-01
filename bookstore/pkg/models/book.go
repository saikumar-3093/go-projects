package models

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/saikumar-3093/go-projects/bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:"unique" json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	err := config.Connect()
	if err != nil {
		fmt.Errorf("Error connecting DB: ", err)
		return
	}
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() *Book {
	if db.NewRecord(b) {
		db.Create(b)
	}
	return b
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookbyID(id int64) (*Book, *gorm.DB) {
	var book Book
	db := db.Where("ID = ?", id).Find(&book)
	return &book, db
}

func DeleteBookbyID(id int64) Book {
	var book Book
	db.Where("ID = ?", id).Delete(book)
	return book
}

// func UpdateBookbyID(id int64, newBook *Book) (*Book, *gorm.DB) {
// 	db := db.Where("ID = ?", id).Update(newBook)
// 	return newBook, db
// }
