package main

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

type LearningTopicResolver struct{}

func (r *LearningTopicResolver) Id(ctx context.Context) *graphql.ID {
	return gqlIDP(1)
}

func (r *LearningTopicResolver) Order(ctx context.Context) int32 {
	return 1
}

func (r *LearningTopicResolver) Description(ctx context.Context) *string {
	data := "stub description"
	return &data
}
