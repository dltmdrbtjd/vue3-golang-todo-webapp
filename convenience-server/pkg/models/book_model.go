package models

import "github.com/graph-gophers/graphql-go"

type Book struct {
	OwnerID uint
	Name    string
	Tags    []Tag
}

type BookInput struct {
	ID      *graphql.ID
	OwnerID int32
	Name    string
	TagIDs  *[]*int32
}
