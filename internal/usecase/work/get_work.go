package work_usecase

import (
	"context"
)

func (uc *workUseCase) FindByID(ctx context.Context, id string) (*WorkOutput, error) {
	workEntity, err := uc.repo.GetByID(ctx, id)

	return &WorkOutput{
		ID:     workEntity.ID,
		Status: string(workEntity.Status),
	}, err
}
