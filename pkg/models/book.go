package models

import (
	"github.com/jinzhu/gorm"
	"github.com/papichulo-coder/go-bookstore/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	//ID  uint64  `gorm:"primary_key;auto_increment" json:"id"`
	// Name   string `json:"name"`
	Name string `gorm:"" json:"name"`
	Type   string `json:"type"`
	Author string `json:"author"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// this is the  create book function where you will create the book
//  b *Books mean you are creating a type of books and  *Books mean you ra returining
// a type book

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b

}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id)
	return &getBook, db
}
func DeleteBook(Id int64) (Book) {
	var book Book
	db.Where("ID=?", Id).Delete(book)
	return book
}
