package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

type DB struct {
	db *gorm.DB
}

func NewDb(path string) (*DB, error) {
	db, err := gorm.Open("sqlite3", path)
	if err != nil {
		return nil, err
	}
	db.DropTableIfExists(&Study{})
	db.AutoMigrate(&Study{})
	for _, v := range studies {
		if err := db.Create(&v).Error; err != nil {
			return nil, err
		}
	}
	return &DB{db}, nil
}

var studies = []Study{
	Study{
		ScopeDefinition:   "Aprender graphql em go",
		SuccessDefinition: "Consigo implementar uma interface graphql em um servidor desenvolvido em go",
	},
}
