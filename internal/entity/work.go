package entity

import (
	"context"
	"time"
)

type WorkStatus string
type WorkType string

// Status works
const (
	WorkStatusPending    WorkStatus = "pending"
	WorkStatusProcessing WorkStatus = "processing"
	WorkStatusCompleted  WorkStatus = "completed"
	WorkStatusFailed     WorkStatus = "failed"
)

// Type works
const (
	WorkTypeAnswer     WorkType = "answer"
	WorkTypeExtract    WorkType = "extract"
	WorkTypeClassify   WorkType = "classify"
	WorkTypeRetrieve   WorkType = "retrieve"
	WorkTypeToolAction WorkType = "tool_action"
)

type Work struct {
	ID             string
	Type           WorkType
	Status         WorkStatus
	ConversationID *string
	Input          []byte
	Output         []byte
	ErrorMessage   *string
	CreatedAt      time.Time
	UpdatedAt      time.Time
	CompletedAt    *time.Time
}

func NewWorkEntity() *Work {
	return &Work{}
}

type WorkRepository interface {
	Create(ctx context.Context, work *Work) error
	GetByID(ctx context.Context, id string) (*Work, error)
}
