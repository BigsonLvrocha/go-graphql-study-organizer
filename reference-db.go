package main

import (
	"github.com/jinzhu/gorm"
)

type Reference struct {
	gorm.Model
	Url        string
	StudyID    uint `gorm:"not null"`
	CategoryID uint `gorm:"not null"`
	Category   ReferenceCategory
}
