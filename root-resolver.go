package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type Resolver struct {
	db *DB
}

type StudyInput struct {
	ScopeDefinition   string
	SuccessDefinition string
}

func (r *Resolver) GetStudy(ctx context.Context, args struct{ ID graphql.ID }) (*StudyResolver, error) {
	id, err := gqlIDToUint(args.ID)
	if err != nil {
		return nil, err
	}
	study, err := r.db.getStudy(ctx, id)
	if err != nil {
		return nil, err
	}
	s := StudyResolver{
		model: *study,
		db:    r.db,
	}
	return &s, nil
}

func (r *Resolver) AddStudy(ctx context.Context, args struct{ Study StudyInput }) (*StudyResolver, error) {
	study, err := r.db.addStudy(ctx, args.Study)
	if err != nil {
		return nil, err
	}
	return &StudyResolver{
		model: *study,
		db:    r.db,
	}, nil
}
