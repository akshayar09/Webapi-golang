package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:abc@tcp(127.0.0.1:3306)/restapi?charset=utf8&parseTime=True")
	if err != nil {
		log.Println("Connection Failed to Open")

	} else {
		log.Println("Connection Established")
	}
	db = d
}

func GetDB() *gorm.DB {
	return db
}
