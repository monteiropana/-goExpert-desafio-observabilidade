package graph

import (
	"github.com/desafio/clean-arch/internal/usecase"
	//go:generate go run github.com/99designs/gqlgen generate
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrderUseCase   usecase.ListOrdersUseCase
}
