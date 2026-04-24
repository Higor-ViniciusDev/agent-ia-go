package work_usecase

import (
	"context"
	"encoding/json"

	entity "github.com/Higor-ViniciusDev/agent-ia-go/internal/domain/work"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/uuid_pkg"
)

type workUseCase struct {
	repo entity.WorkRepositoryInterface
}

func New(repo entity.WorkRepositoryInterface) *workUseCase {
	return &workUseCase{repo: repo}
}

func (uc *workUseCase) Execute(ctx context.Context, input WorkInput) (*WorkOutput, error) {
	work := entity.NewWorkEntity()

	work.Type = entity.WorkType(input.Type)
	work.ID = uuid_pkg.NewID().String()
	work.Status = entity.WorkStatusPending
	work.ConversationID = input.ConversationID

	// serializa o Data para JSONB
	encoded, err := json.Marshal(input.Data)
	if err != nil {
		return nil, err
	}
	work.Input = encoded

	if err := uc.repo.Create(ctx, work); err != nil {
		return nil, err
	}

	return &WorkOutput{
		ID:     work.ID,
		Status: string(work.Status),
	}, nil
}
