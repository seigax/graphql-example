package graph

import (
	"errors"

	"github.com/seigax/graphql-example/graph/model"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoLists []*model.Todo
	Users     []*model.User
}

func (r *Resolver) DoneTodo(input model.DoneTodo) (*model.Todo, error) {
	for _, v := range r.TodoLists {
		if v.ID == input.ID {
			v.Done = true
			return v, nil
		}
	}
	return nil, errors.New("Data Not Found")
}

func (r *Resolver) UndoneTodo(input model.UnDoneTodo) (*model.Todo, error) {
	for _, v := range r.TodoLists {
		if v.ID == input.ID {
			v.Done = false
			return v, nil
		}
	}
	return nil, errors.New("Data Not Found")
}
