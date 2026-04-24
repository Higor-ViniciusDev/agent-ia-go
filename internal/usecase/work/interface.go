package work_usecase

import "context"

type WorkUseCaseInterface interface {
	Execute(ctx context.Context, input WorkInput) (*WorkOutput, error)
	FindByID(ctx context.Context, id string) (*WorkOutput, error)
}
