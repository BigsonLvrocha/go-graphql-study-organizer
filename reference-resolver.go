package main

import "context"

type ReferenceResolver struct{}

func (r *ReferenceResolver) Url(ctx context.Context) *string {
	data := "açsldkfjaçslkdf.com"
	return &data
}

func (r *ReferenceResolver) Category(ctx context.Context) *ReferenceCategoryResolver {
	data := ReferenceCategoryResolver{}
	return &data
}
