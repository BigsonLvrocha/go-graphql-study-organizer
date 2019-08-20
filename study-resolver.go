package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type StudyResolver struct {
	model *Study
}

func (s *StudyResolver) Id(ctx context.Context) *graphql.ID {
	return gqlIDP(1)
}

func (s *StudyResolver) ScopeDefinition(ctx context.Context) *string {
	return &s.model.ScopeDefinition
}

func (s *StudyResolver) SuccessDefinition(ctx context.Context) *string {
	return &s.model.SuccessDefinition
}

func (s *StudyResolver) LearningPath(ctx context.Context) *[]*LearningTopicResolver {
	data := make([]*LearningTopicResolver, 1)
	data[0] = &LearningTopicResolver{}
	return &data
}

func (s *StudyResolver) References(ctx context.Context) *[]*ReferenceResolver {
	data := make([]*ReferenceResolver, 1)
	data[0] = &ReferenceResolver{}
	return &data
}
