package model

import (
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func SetUpDatabase(dbpath string) {
	db, err := gorm.Open(sqlite.Open(dbpath), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		panic("failed to connect database")
	}
	DB = db
}
