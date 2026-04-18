package entity

import "context"

type WorkRepositoryInterface interface {
	Create(ctx context.Context, work *Work) error
	GetByID(ctx context.Context, id string) (*Work, error)
}
