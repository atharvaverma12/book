package models

import {
	"github.com/jinzhu/gorm"
	"github.com/atharvaverma12/book/pkg/config"
}

//use of orm

var db * gorm.db

type Book struct{
	gorm.Model
	Name string `gorm:""json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

//perform the logic here

func (b * Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books [] Book 
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB){
	var getBook Book
	db := db.Where("ID=?",Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int 64) Book {
	var book Book
    db.Where("ID=?", ID).Delete(book)
	return book
}