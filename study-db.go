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
	References        []Reference
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

func (db *DB) getStudyReferences(ctx context.Context, studyID uint) ([]Reference, error) {
	var s Study
	var refs []Reference

	s.ID = studyID
	err := db.db.Model(&s).Association("References").Find(&refs).Error
	if err != nil {
		return nil, err
	}
	return refs, nil
}
