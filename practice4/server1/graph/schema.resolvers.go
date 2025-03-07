package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"practice4/db"
	"practice4/graph/model"
)

// GetAllProducts is the resolver for the GetAllProducts field.
func (r *queryResolver) GetAllProducts(ctx context.Context) ([]*model.Product, error) {
	products, err := db.GetAllProducts(r.DB)
	if err != nil {
		return nil, err
	}

	return products, nil
}

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
