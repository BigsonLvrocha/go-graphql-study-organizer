package main

import (
	"github.com/jinzhu/gorm"
)

type LearningTopic struct {
	gorm.Model
	Order       int32 `gorm:"not null"`
	Description string
	StudyID     uint
}
