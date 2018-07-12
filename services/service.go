package services

import (
	"context"
)

// This is a general service interface. Since usually services are just one of
// the kind of CRUD.
type Service interface {
	Retrieve(ctx context.Context, args interface{}) (interface{}, error)

	Create(ctx context.Context, args interface{}) (interface{}, error)

	Update(ctx context.Context, args interface{}) (interface{}, error)

	Delete(ctx context.Context, args interface{}) (interface{}, error)
}
