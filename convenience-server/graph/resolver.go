package graph

//go:generate go run github.com/99designs/gqlgen generate
import "github.com/Convenience-Tools/convenience-server/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.
// script : go generate ./...
type Resolver struct {
	todos []*model.Todo
}
