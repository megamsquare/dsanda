package controllers

import (
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/megamsquare/dsanda/models"
	"github.com/megamsquare/dsanda/utils"
)

var NewBook models.Book

func CreateBook(w http.ResponseWriter, r *http.Request) {
	CreateBook := &models.Book{}
	err := utils.ParseBody(r, CreateBook)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	CreateBook, err = CreateBook.CreateBook()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusCreated, CreateBook)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books, err := NewBook.GetBooks()
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, books)
}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	id, _ := strconv.ParseInt(bookId, 0, 0)
	book, err := NewBook.GetBook(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	UpdateBook := &models.Book{}
	err := utils.ParseBody(r, UpdateBook)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	params := mux.Vars(r)
	bookId := params["id"]
	id, _ := strconv.ParseInt(bookId, 0, 0)
	UpdateBook, err = UpdateBook.UpdateBook(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, UpdateBook)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookId := params["id"]
	id, _ := strconv.ParseInt(bookId, 0, 0)
	_, err := NewBook.DeleteBook(id)
	if err != nil {
		utils.RespondWithError(w, http.StatusBadRequest, "Invalid Book ID")
		return
	}
	utils.RespondWithJSON(w, http.StatusOK, map[string]string{"result": "success"})
}
