package repository

import (
	"context"
	"database/sql"

	entity "github.com/Higor-ViniciusDev/agent-ia-go/internal/domain/work"
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
	return nil
}

func (w *WorkRepository) GetByID(ctx context.Context, id string) (*entity.Work, error) {
	return nil, nil
}
