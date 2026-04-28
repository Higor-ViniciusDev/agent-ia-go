package work_usecase

import (
	"context"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/internal_error"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/uuid_pkg"
)

func (uc *workUseCase) FindByID(ctx context.Context, id string) (*WorkOutput, *internal_error.InternalError) {

	if _, err := uuid_pkg.PaserID(id); err != nil {
		return nil, internal_error.NewBadRequestError("uuid invalid")
	}

	workEntity, err := uc.repo.GetByID(ctx, id)

	if err != nil {
		return nil, err
	}

	return &WorkOutput{
		ID:     workEntity.ID,
		Status: string(workEntity.Status),
	}, err
}
