package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type Resolver struct {
	db *DB
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
		model: study,
	}
	return &s, nil
}
