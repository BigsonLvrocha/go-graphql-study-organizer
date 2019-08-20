package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type Resolver struct{}

func (r *Resolver) GetStudy(ctx context.Context, args struct{ ID graphql.ID }) *StudyResolver {
	return &StudyResolver{}
}
