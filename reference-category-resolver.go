package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type ReferenceCategoryResolver struct{}

func (r *ReferenceCategoryResolver) Id(ctx context.Context) *graphql.ID {
	return gqlIDP(1)
}

func (r *ReferenceCategoryResolver) Name(ctx context.Context) *string {
	data := "stub category"
	return &data
}
