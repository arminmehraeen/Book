package models

import (
	"books/pkg/config"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	// ID          uint       `gorm:"primary_key" json:"id"`
	gorm.Model
	Name        string     `json:"name"`
	Author      string     `json:"author"`
	Publication string     `json:"publication"`
	// CreatedAt   time.Time  `json:"created_at"`
	// UpdatedAt   time.Time  `json:"updated_at"`
	// DeletedAt   *time.Time `json:"deleted_at"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func Index() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func (b *Book) Create() *Book {
	id := db.Create(&Book{
		Name: "Armin",
		Author: "Naser",
		Publication: "Ebi",
	})
	fmt.Println(id)
	return b
}

func Show(Id int64) (*Book , *gorm.DB) {
	var book Book
	db := db.Where("id=?",Id).Find(&book)
	return &book , db
}

func Delete(Id int64) Book{
	var book Book
	db.Where("id=?",Id).Delete(&book)
	return book 
}
