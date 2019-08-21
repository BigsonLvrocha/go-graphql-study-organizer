package main

import "context"

type ReferenceResolver struct {
	db    *DB
	model Reference
}

func (r *ReferenceResolver) Url(ctx context.Context) *string {
	return &r.model.Url
}

func (r *ReferenceResolver) Category(ctx context.Context) (*ReferenceCategoryResolver, error) {
	cat, err := r.db.getReferenceCategory(ctx, r.model.CategoryID)
	if err != nil {
		return nil, err
	}
	return &ReferenceCategoryResolver{
		model: *cat,
	}, nil
}
