package main

import (
	"context"

	"github.com/jinzhu/gorm"
)

type Study struct {
	gorm.Model
	ScopeDefinition   string
	SuccessDefinition string
}

func (db *DB) getStudy(ctx context.Context, studyID uint) (*Study, error) {
	var User Study
	err := db.db.First(&User, studyID).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}
