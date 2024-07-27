package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/skanenje/11-projects/go-bookstore/pkg/utils"
	"github.com/skanenje/11-projects/go-bookstore/pkg/models"
	// "github.com/skanenje/go-bookstore/pkg/utils"
)
var NewBook models.Book

func GetBook(w http.ResponseWriter, r *http.Request){
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0 ,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	bookDeatails, _ := models.GetBookById(Id)
	res, _ := json.Marshal(bookDeatails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.Request){
	CreateBook := &models.Book{}
	utils.ParseBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteBook(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0 ,0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	book := models.DeleteBook(Id)
	res, _ := json.Marshal(book)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook = &models.Book{}
	utils.ParseBody(r, updateBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	id , err := strconv.ParseInt(bookId, 0 , 0)
	if err != nil {
		fmt.Println("error parsing")
	}
	bookdetails, db := models.GetBookById(id)
	if updateBook.Name != ""{
		bookdetails.Name = updateBook.Name
	}
	if updateBook.Author != ""{
		bookdetails.Author = updateBook.Author
	}
	if updateBook.Publication != ""{
		bookdetails.Publication = updateBook.Publication
	}
	db.Save(&bookdetails)
	res, _ := json.Marshal(bookdetails)
	w.Header().Set("content-type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}