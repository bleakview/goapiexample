package model

import (
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

type Book struct {
	ID           uuid.UUID `gorm:"primary_key" json:"id"`
	Name         string    `json:"name"`
	Author       string    ` json:"author"`
	Release_year int       `json:"release_year"`
	ISBN         string    `json:"isbn"`
}

func (book *Book) BeforeCreate(db *gorm.DB) error {
	if book.ID.String() == "00000000-0000-0000-0000-000000000000" {
		uuid := uuid.NewV4().String()
		db.Statement.SetColumn("ID", uuid)
	}
	return nil
}

func DBMigrate(db *gorm.DB) *gorm.DB {
	db.AutoMigrate(&Book{})
	return db
}
