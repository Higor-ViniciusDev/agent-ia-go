package work_usecase

import (
	"context"
	"encoding/json"

	entity "github.com/Higor-ViniciusDev/agent-ia-go/internal/domain/work"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/internal_error"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/events"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/logger"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/uuid_pkg"
)

type workUseCase struct {
	repo             entity.WorkRepositoryInterface
	eventWorkCreated events.EventInterface
	eventDispachet   events.EventDispachetInterface
}

func New(repo entity.WorkRepositoryInterface, eventWorkCreated events.EventInterface, eventDis events.EventDispachetInterface) *workUseCase {
	return &workUseCase{
		repo:             repo,
		eventWorkCreated: eventWorkCreated,
		eventDispachet:   eventDis,
	}
}

func (uc *workUseCase) Execute(ctx context.Context, input WorkInput) (*WorkOutput, *internal_error.InternalError) {
	work := entity.NewWorkEntity()

	work.Type = entity.WorkType(input.Type)
	work.ID = uuid_pkg.NewID().String()
	work.Status = entity.WorkStatusPending
	work.ConversationID = input.ConversationID

	// serializa o Data para JSONB
	encoded, err := json.Marshal(input.Data)
	if err != nil {
		return nil, internal_error.NewInternalServerError("Error convert json for byte")
	}
	work.Input = encoded

	if err := uc.repo.Create(ctx, work); err != nil {
		return nil, err
	}

	response := &WorkOutput{
		ID:     work.ID,
		Status: string(work.Status),
	}

	uc.eventWorkCreated.SetValues(response)
	errDispatch := uc.eventDispachet.Dispatch(uc.eventWorkCreated)

	if errDispatch != nil {
		logger.Error("Error in dispachter workCreated: ", errDispatch)
	}

	return response, nil
}
