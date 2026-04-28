package repository

import (
	"context"
	"database/sql"

	entity "github.com/Higor-ViniciusDev/agent-ia-go/internal/domain/work"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/internal_error"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/logger"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/uuid_pkg"
)

type WorkRepository struct {
	db *sql.DB
}

func NewWorkRepository(Db *sql.DB) *WorkRepository {
	return &WorkRepository{
		db: Db,
	}
}

func (w *WorkRepository) Create(ctx context.Context, work *entity.Work) *internal_error.InternalError {
	query := `INSERT INTO works (id, type, status, conversation_id, input, output, error_message) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := w.db.ExecContext(ctx, query, work.ID, work.Type, work.Status, work.ConversationID, work.Input, work.Output, work.ErrorMessage)

	if err != nil {
		logger.Error("Error creating work: ", err)
		return internal_error.NewInternalServerError("Error creating user")
	}

	return nil
}

func (w *WorkRepository) GetByID(ctx context.Context, id string) (*entity.Work, *internal_error.InternalError) {
	query := `SELECT id, type, status, conversation_id, input, output, 
	error_message FROM works WHERE id = $1`
	row := w.db.QueryRowContext(ctx, query, id)

	if _, err := uuid_pkg.PaserID(id); err != nil {
		return nil, internal_error.NewBadRequestError("uuid invalid")
	}

	work := entity.NewWorkEntity()
	if err := row.Scan(&work.ID, &work.Type, &work.Status, &work.ConversationID, &work.Input, &work.Output, &work.ErrorMessage); err != nil {
		if err == sql.ErrNoRows {
			return nil, internal_error.NewBadRequestError("Work not found")
		}

		logger.Error("Error finding work by ID: ", err)
		return nil, internal_error.NewInternalServerError("Error finding user")
	}

	return work, nil
}
