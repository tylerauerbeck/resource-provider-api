package api

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.32

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// ResourceProvider returns ResourceProviderResolver implementation.
func (r *Resolver) ResourceProvider() ResourceProviderResolver { return &resourceProviderResolver{r} }

type queryResolver struct{ *Resolver }
type resourceProviderResolver struct{ *Resolver }