package initdb

import (
	"encoding/json"

	"github.com/bleakview/goapiexample/cmd/goapiexample/model"
	"gorm.io/gorm"
)

//generate test data
func InitDB(db *gorm.DB) {
	CleanDB(db)
	var book model.Book
	sampleDbData := `{
		"id": "5d596c01-e20b-4049-91e9-a0be77715260",
		"name": "name",
		"author": "author",
		"release_year": 1980,
		"isbn": "isbn"
	  }`
	json.Unmarshal([]byte(sampleDbData), &book)

	db.Save(&book)

}
func CleanDB(db *gorm.DB) {
	db.Where("1=1").Delete(model.Book{})

}
