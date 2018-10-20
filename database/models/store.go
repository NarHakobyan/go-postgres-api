package models

import (
	. "github.com/NarHakobyan/go-postgres-api/database"
)

//go:generate goqueryset -in store.go

// Store struct represents user model.
// gen:qs
type Store struct {
	Model
	Name     string `form:"name" json:"name" valid:"required~Name is required"`
	Capacity int    `form:"capacity" json:"capacity" valid:"required~Capacity is required"`
}

//var StoreQuery = NewStoreQuerySet(Db)
