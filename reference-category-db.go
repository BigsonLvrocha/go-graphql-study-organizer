package main

import (
	"context"

	"github.com/jinzhu/gorm"
)

type ReferenceCategory struct {
	gorm.Model
	Name string
}

func (db *DB) getReferenceCategory(ctx context.Context, referenceCategoryID uint) (*ReferenceCategory, error) {
	var cat ReferenceCategory
	err := db.db.First(&cat, referenceCategoryID).Error
	if err != nil {
		return nil, err
	}
	return &cat, nil
}
