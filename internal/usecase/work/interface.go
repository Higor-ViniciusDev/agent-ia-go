package work_usecase

import (
	"context"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/internal_error"
)

type WorkUseCaseInterface interface {
	Execute(ctx context.Context, input WorkInput) (*WorkOutput, *internal_error.InternalError)
	FindByID(ctx context.Context, id string) (*WorkOutput, *internal_error.InternalError)
}
