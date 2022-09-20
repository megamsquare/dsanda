package models

import (
	"gorm.io/gorm"
	"github.com/megamsquare/dsanda/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Title  string `gorm:"not null" json:"title"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (b *Book) CreateBook() (*Book, error) {
	if err := db.Create(&b).Error; err != nil {
		return &Book{}, err
	}
	return b, nil
}

func (b *Book) GetBooks() (*[]Book, error) {
	books := []Book{}
	err := db.Find(&books).Error
	if err != nil {
		return &[]Book{}, err
	}
	return &books, err
}

func (b *Book) GetBook(id int64) (*Book, error) {
	var book Book
	err := db.Where("id = ?", id).Find(&book).Error
	if err != nil {
		return &Book{}, err
	}
	return &book, nil
}

func (b *Book) UpdateBook(id int64) (*Book, error) {
	bookDetails, err := b.GetBook(id)
	if err != nil {
		return &Book{}, err
	}
	if b.Title != "" {
		bookDetails.Title = b.Title
	}
	if b.Author != "" {
		bookDetails.Author = b.Author
	}
	if b.Publication != "" {
		bookDetails.Publication = b.Publication
	}
	err = db.Save(&bookDetails).Error
	if err != nil {
		return &Book{}, err
	}
	return bookDetails, nil
}

func (b *Book) DeleteBook(id int64) (int64, error) {
	var book Book
	err := db.Where("id = ?", id).Delete(&book).Error
	if err != nil {
		return 0, err
	}
	return db.RowsAffected, nil
}