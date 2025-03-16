package egq

import (
	"context"

	"gorm.io/gorm"
)

type Repository[T any] interface {
	Create(context.Context, *T) error
	UpdateByID(context.Context, uint, *T) error
	DeleteByID(context.Context, uint) error

	Find(context.Context, Request) (*Response, error)
	FindOne(context.Context, Request) (*Response, error)
}

type BaseRepository[T any] struct {
	DB           *gorm.DB
	Model        T
	SearchFields []string
	AllowedSorts []string
	Validations  map[string]func(interface{}) error
}
