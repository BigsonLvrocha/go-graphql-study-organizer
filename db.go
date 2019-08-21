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
	db.DropTableIfExists(&Study{}, &LearningTopic{}, &Reference{})
	db.AutoMigrate(&Study{}, &LearningTopic{}, &Reference{})
	for _, v := range studies {
		if err := db.Create(&v).Error; err != nil {
			return nil, err
		}
	}
	for _, v := range learningTopics {
		if err := db.Create(&v).Error; err != nil {
			return nil, err
		}
	}
	for _, v := range references {
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

var learningTopics = []LearningTopic{
	LearningTopic{
		StudyID:     1,
		Order:       1,
		Description: "Aprender o básico da sintaxe do go",
	},
	LearningTopic{
		StudyID:     1,
		Order:       2,
		Description: "Aprender a usar Go routines",
	},
	LearningTopic{
		StudyID:     1,
		Order:       3,
		Description: "Aprender a utilizar o gerenciamento de dependências",
	},
}

var references = []Reference{
	Reference{
		StudyID: 1,
		Url:     "https://golang.org/",
	},
	Reference{
		StudyID: 1,
		Url:     "https://golang.org/doc/code.html?h=tour",
	},
}
