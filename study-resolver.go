package main

import (
	"context"

	graphql "github.com/graph-gophers/graphql-go"
)

type StudyResolver struct {
	db    *DB
	model Study
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

func (s *StudyResolver) LearningPath(ctx context.Context) (*[]*LearningTopicResolver, error) {
	lts, err := s.db.getStudyLearningTopics(ctx, s.model.ID)
	if err != nil {
		return nil, err
	}
	data := make([]*LearningTopicResolver, len(lts))
	for i, lt := range lts {
		data[i] = &LearningTopicResolver{
			model: lt,
		}
	}
	return &data, nil
}

func (s *StudyResolver) References(ctx context.Context) *[]*ReferenceResolver {
	data := make([]*ReferenceResolver, 1)
	data[0] = &ReferenceResolver{}
	return &data
}
