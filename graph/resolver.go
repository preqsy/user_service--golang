package graph

import "user_service/core"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
//go:generate go run github.com/99designs/gqlgen generate

type Resolver struct {
	service *core.Service
}

func NewResolver(service *core.Service) *Resolver {
	return &Resolver{service: service}
}
