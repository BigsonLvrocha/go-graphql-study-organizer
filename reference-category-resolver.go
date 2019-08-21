package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type ReferenceCategoryResolver struct {
	model ReferenceCategory
}

func (r *ReferenceCategoryResolver) ID(ctx context.Context) *graphql.ID {
	return gqlIDP(r.model.ID)
}

func (r *ReferenceCategoryResolver) Name(ctx context.Context) *string {
	return &r.model.Name
}
