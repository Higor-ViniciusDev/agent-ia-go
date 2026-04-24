package service

import (
	"context"
	"encoding/json"

	"github.com/Higor-ViniciusDev/agent-ia-go/internal/infra/grpc/proto/pb"
	work_usecase "github.com/Higor-ViniciusDev/agent-ia-go/internal/usecase/work"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type WorkService struct {
	useCase work_usecase.WorkUseCaseInterface
	pb.UnimplementedWorkServiceServer
}

func NewWorkService(workUsecase work_usecase.WorkUseCaseInterface) *WorkService {
	return &WorkService{
		useCase: workUsecase,
	}
}

func (w *WorkService) WorkAction(ctx context.Context, req *pb.WorkRequest) (*pb.ResponseWorkAction, error) {
	if req.Data == nil {
		return nil, status.Error(codes.InvalidArgument, "data is required")
	}

	// google.protobuf.Struct → map[string]any
	dataMap := req.Data.AsMap()

	workType, _ := dataMap["type"].(string)
	if workType == "" {
		return nil, status.Error(codes.InvalidArgument, "data.type is required")
	}

	var conversationID *string
	if v, ok := dataMap["conversation_id"].(string); ok && v != "" {
		conversationID = &v
	}

	input := work_usecase.WorkInput{
		Type:           workType,
		ConversationID: conversationID,
		Data:           dataMap,
	}

	workOutput, err := w.useCase.Execute(ctx, input)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	jsonBytes, err := json.Marshal(workOutput)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &pb.ResponseWorkAction{
		Response: string(jsonBytes),
	}, nil
}

func (c *WorkService) GetWorkById(ctx context.Context, req *pb.GetWorkByIdInput) (*pb.Work, error) {
	work, err := c.useCase.FindByID(ctx, req.Id)

	if err != nil {
		return nil, err
	}

	return &pb.Work{
		Id:     work.ID,
		Status: work.Status,
	}, nil
}
