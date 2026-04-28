package repository

import (
	"context"
	"database/sql"

	entity "github.com/Higor-ViniciusDev/agent-ia-go/internal/domain/work"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/internal_error"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/logger"
)

type WorkRepository struct {
	db *sql.DB
}

func NewWorkRepository(Db *sql.DB) *WorkRepository {
	return &WorkRepository{
		db: Db,
	}
}

func (w *WorkRepository) Create(ctx context.Context, work *entity.Work) error {
	query := `INSERT INTO works (id, type, status, conversation_id, input, output, error_message) VALUES ($1, $2, $3, $4, $5, $6, $7)`
	_, err := w.db.ExecContext(ctx, query, work.ID, work.Type, work.Status, work.ConversationID, work.Input, work.Output, work.ErrorMessage)

	if err != nil {
		logger.Error("Error creating work: ", err)
		return internal_error.NewInternalServerError("Error creating user")
	}

	return nil
}

func (w *WorkRepository) GetByID(ctx context.Context, id string) (*entity.Work, error) {
	query := `SELECT id, type, status, conversation_id, input, output, 
	error_message FROM works WHERE id = $1`
	row := w.db.QueryRowContext(ctx, query, id)

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
