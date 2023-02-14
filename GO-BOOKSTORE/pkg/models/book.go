package models

import (
	"GO-BOOKSTORE/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})

	db.Create(&Book{Name: "john", Author: "doe", Publication: "ab"})
	db.Create(&Book{Name: "jane", Author: "smith", Publication: "cd"})
	db.Create(&Book{Name: "mary", Author: "philip", Publication: "ef"})
	db.Create(&Book{Name: "emma", Author: "swan", Publication: "gh"})
	db.Create(&Book{Name: "meredith", Author: "grey", Publication: "ij"})
	db.Create(&Book{Name: "rachel", Author: "green", Publication: "kl"})
	db.Create(&Book{Name: "monica", Author: "geller", Publication: "mn"})

}

func (b *Book) CreateBook() *Book {
	db.NewRecord(b)
	db.Create(&b)
	return b
}

func GetAllBooks() []Book {
	var Books []Book
	db.Find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", Id).Find(&getBook)
	return &getBook, db
}

func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}
