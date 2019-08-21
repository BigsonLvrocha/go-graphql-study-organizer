package main

import (
	"github.com/jinzhu/gorm"
)

type Reference struct {
	gorm.Model
	Url     string
	StudyID uint
}
