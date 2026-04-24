package work_usecase

import (
	"context"

	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/logger"
)

func (uc *workUseCase) FindByID(ctx context.Context, id string) (*WorkOutput, error) {
	workEntity, err := uc.repo.GetByID(ctx, id)

	logger.Info(id)
	return &WorkOutput{
		ID:     workEntity.ID,
		Status: string(workEntity.Status),
	}, err
}
