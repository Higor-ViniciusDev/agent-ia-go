package repository

import (
	"context"
	"database/sql"
	"encoding/json"
	"os"
	"testing"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/config"
	entity "github.com/Higor-ViniciusDev/agent-ia-go/internal/domain/work"
	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/database"
	"github.com/Higor-ViniciusDev/agent-ia-go/pkg/uuid_pkg"
)

func TestCreateNotWithIdConversation(t *testing.T) {
	db := mockDB()

	if db == nil {
		t.Fatalf("Fail: empty db connection")
	}
	work_test := mockNewWork()

	ctx := context.Background()
	repository := NewWorkRepository(db)

	// serializa o Data para JSONB
	encoded, errJson := json.Marshal("{teste:1}")
	if errJson != nil {
		t.Errorf("Error convert json for bytes date")
	}

	work_test.Input = encoded
	defer repository.DeleteAllWorks(ctx)
	err := repository.Create(ctx, work_test)

	if err != nil {
		t.Errorf("Fail in save date repository")
	}

	dateWork, err := repository.GetByID(ctx, work_test.ID)

	if err != nil {
		t.Errorf("Fail in search date repository")
	}

	if dateWork.ID != work_test.ID || dateWork.Type != work_test.Type {
		t.Errorf("date work differ to the original")
	}
}

func TestGetWorkByIdNotFound(t *testing.T) {
	db := mockDB()
	ctx := context.Background()
	repository := NewWorkRepository(db)

	fakeID := uuid_pkg.NewID().String()
	_, err := repository.GetByID(ctx, fakeID)

	if err == nil {
		t.Errorf("expected error for non-existent ID, got nil")
	}
}

func mockDB() *sql.DB {
	os.Setenv("DB_HOST", "localhost")
	os.Setenv("DB_PORT", "5432")
	os.Setenv("DB_USER", "postgres")
	os.Setenv("DB_PASSWORD", "postgres")
	os.Setenv("DB_NAME", "work_agent")

	config := config.Load()

	newConexao := database.NewConnect(config)

	return newConexao
}

func mockNewWork() *entity.Work {
	work := entity.NewWorkEntity()

	work.Type = entity.WorkType("answer")
	work.ID = uuid_pkg.NewID().String()
	work.Status = entity.WorkStatusPending
	work.ConversationID = nil

	return work
}
