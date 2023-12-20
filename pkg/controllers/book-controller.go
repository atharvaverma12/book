package controllers
import {
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/atharvaverma12/book/pkg/utils"
	"github.com/atharvaverma12/book/pkg/models"
}

var NewBook models.Book

func GetBook(w http.ResponseWriter, r * http.Request){
	NewBooks := models.GetAllBooks()
	res, _ := josn.Marshal(NewBooks)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r * http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err!= nil{
		fmt.Println("error while parsing")
	}
	bookDetails, _:= models.GetBookById(ID)
	res,_ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBook(w http.ResponseWriter, r * http.Request){
	CreateBook := &models.Book{}
	Utils.ParseBody(r,CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b) //convert it into json from getting by db
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r * http.Request){
	vars := mux.Vars(r)
	bookId := vars["BookID"]
	ID , err := strconv.ParseInt(bookId,0,0)
	if err!= nil{
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(ID)
	res,_ := json.Marshal(book)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r * http.Request){
	//delete old book
	//create new book
	var UpdateBook = &models.Book{}
	utils.ParseBody(r,UpdateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId,0,0)
	if err!= nil {
		fmt.Println("error while parsing")
	}

	bookDetails, db := models.GetBookById(ID)

	if UpdateBook.Name != ""{
		bookDetails.Name = UpdateBook.Name
	}
	if UpdateBook.Author != ""{
		bookDetails.Author = UpdateBook.Author
	}
	if UpdateBook.Publication != ""{
		bookDetails.Publication = UpdateBook.Publication
	}
	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}