package main

import "context"

type ReferenceResolver struct {
	model Reference
}

func (r *ReferenceResolver) Url(ctx context.Context) *string {
	return &r.model.Url
}

func (r *ReferenceResolver) Category(ctx context.Context) *ReferenceCategoryResolver {
	data := ReferenceCategoryResolver{}
	return &data
}
