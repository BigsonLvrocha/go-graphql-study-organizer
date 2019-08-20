package main

import (
	"context"

	"github.com/jinzhu/gorm"
)

type Study struct {
	gorm.Model
	ScopeDefinition   string
	SuccessDefinition string
	LearningTopics    []LearningTopic
}

func (db *DB) getStudy(ctx context.Context, studyID uint) (*Study, error) {
	var User Study
	err := db.db.First(&User, studyID).Error
	if err != nil {
		return nil, err
	}
	return &User, nil
}

func (db *DB) getStudyLearningTopics(ctx context.Context, studyID uint) ([]LearningTopic, error) {
	var s Study
	var lt []LearningTopic

	s.ID = studyID
	err := db.db.Model(&s).Association("LearningTopics").Find(&lt).Error
	if err != nil {
		return nil, err
	}
	return lt, nil
}
