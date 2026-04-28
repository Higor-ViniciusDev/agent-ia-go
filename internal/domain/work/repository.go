package entity

import (
	"context"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/internal_error"
)

type WorkRepositoryInterface interface {
	Create(ctx context.Context, work *Work) *internal_error.InternalError
	GetByID(ctx context.Context, id string) (*Work, *internal_error.InternalError)
}
