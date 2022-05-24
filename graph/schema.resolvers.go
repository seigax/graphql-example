package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"math/rand"

	"github.com/seigax/graphql-example/graph/generated"
	"github.com/seigax/graphql-example/graph/model"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo := model.Todo{
		ID:   fmt.Sprint("ID_", rand.Int()),
		Text: input.Text,
		Done: false,
		User: &model.User{
			ID:   fmt.Sprint("ID_", rand.Int()),
			Name: fmt.Sprint("NAME_", rand.Int()),
		},
	}
	r.TodoLists = append(r.TodoLists, &todo)
	return &todo, nil
}

func (r *mutationResolver) DoneTodo(ctx context.Context, input model.DoneTodo) (*model.Todo, error) {
	return r.Resolver.DoneTodo(input)
}

func (r *mutationResolver) UndoneTodo(ctx context.Context, input model.UnDoneTodo) (*model.Todo, error) {
	return r.Resolver.UndoneTodo(input)
}

func (r *queryResolver) Todos(ctx context.Context) ([]*model.Todo, error) {
	return r.TodoLists, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
