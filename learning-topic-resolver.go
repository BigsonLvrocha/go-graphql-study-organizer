package main

import (
	"context"

	"github.com/graph-gophers/graphql-go"
)

type LearningTopicResolver struct {
	model LearningTopic
}

func (r *LearningTopicResolver) ID(ctx context.Context) *graphql.ID {
	return gqlIDP(r.model.ID)
}

func (r *LearningTopicResolver) Order(ctx context.Context) int32 {
	return r.model.Order
}

func (r *LearningTopicResolver) Description(ctx context.Context) *string {
	return &r.model.Description
}
