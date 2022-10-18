package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/bostigger/bukstore/pkg/models"
	"github.com/bostigger/bukstore/pkg/utils"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
)

var NewBook models.Book

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	AllStoreBooks := models.GetAllBooks()
	res, _ := json.Marshal(AllStoreBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetStoreBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error during parsing...")
	}
	bookDetails, _ := models.GetBookById(ID)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateNewStoreBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	utils.DigestBody(r, CreateBook)
	b := CreateBook.CreateBook()
	res, _ := json.Marshal(b)
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteStoreBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error during parsing...")
	}
	book := models.DeleteBook(ID)
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateStoreBook(w http.ResponseWriter, r *http.Request) {
	var updateStoreBook = &models.Book{}
	utils.DigestBody(r, updateStoreBook)
	vars := mux.Vars(r)
	bookId := vars["bookId"]
	ID, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println("Error during parsing...")
	}
	bookDetails, db := models.GetBookById(ID)
	if updateStoreBook != nil {
		bookDetails.Name = updateStoreBook.Name

	}
	if updateStoreBook.Author != "" {
		bookDetails.Author = updateStoreBook.Author
	}
	if updateStoreBook.Publication != "" {
		bookDetails.Publication = updateStoreBook.Publication
	}

	db.Save(&bookDetails)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
