package handler

import (
	"encoding/json"
	"net/http"

	"github.com/bleakview/goapiexample/cmd/goapiexample/model"
	"github.com/gorilla/mux"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

func GetAllBooks(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	books := []model.Book{}
	db.Find(&books)
	respondJSON(w, http.StatusOK, books)
}

func CreateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	bookRequest := model.BookRequest{}

	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&bookRequest); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()

	book := model.Book{
		Name:         bookRequest.Name,
		Author:       bookRequest.Author,
		Release_year: bookRequest.Release_year,
		ISBN:         bookRequest.ISBN,
	}

	if err := db.Save(&book).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusCreated, book)
}

func GetBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	book := getBookOrnil(db, id, w, r)
	if book == nil {
		return
	}
	respondJSON(w, http.StatusOK, book)
}

func UpdateBook(db *gorm.DB, w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id := vars["id"]
	book := getBookOrnil(db, id, w, r)
	//do not allow id change with id we will set it later
	bookId := book.ID
	if book == nil {
		return
	}

	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&book); err != nil {
		respondError(w, http.StatusBadRequest, err.Error())
		return
	}
	defer r.Body.Close()
	book.ID = bookId
	if err := db.Save(&book).Error; err != nil {
		respondError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondJSON(w, http.StatusOK, book)
}

func getBookOrnil(db *gorm.DB, id string, w http.ResponseWriter, r *http.Request) *model.Book {
	uuidId, _ := uuid.FromString(id)
	book := model.Book{}
	if err := db.First(&book, model.Book{ID: uuidId}).Error; err != nil {
		respondError(w, http.StatusNotFound, err.Error())
		return nil
	}
	return &book
}
