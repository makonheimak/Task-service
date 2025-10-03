package grpc

import (
	"context"
	"fmt"
	"log"

	"github.com/makonheimak/task-service/internal/task/orm"
	"github.com/makonheimak/task-service/internal/task/service"

	pb "github.com/makonheimak/project-protos/proto/task"
	pbUser "github.com/makonheimak/project-protos/proto/user"

	"google.golang.org/protobuf/types/known/emptypb"
)

type Handler struct {
	svc        *service.Service
	userClient pbUser.UserServiceClient
	pb.UnimplementedTaskServiceServer
}

func NewHandler(svc *service.Service, uc pbUser.UserServiceClient) *Handler {
	return &Handler{svc: svc, userClient: uc}
}

func (h *Handler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.CreateTaskResponse, error) {
	if _, err := h.userClient.GetUserByID(ctx, &pbUser.GetUserByIDRequest{Id: req.UserId}); err != nil {
		return nil, fmt.Errorf("user %d not found: %w", req.UserId, err)
	}

	log.Printf("gRPC CreateTask called with Task=%s", req.Task)

	task, err := h.svc.CreateTask(orm.Task{
		Task:   req.Task,
		UserID: req.UserId,
	})
	if err != nil {
		return nil, err
	}

	return &pb.CreateTaskResponse{
		Task: &pb.Task{
			Id:     task.ID,
			Task:   task.Task,
			UserId: task.UserID,
		},
	}, nil
}

func (h *Handler) GetAllTasks(ctx context.Context, req *pb.GetAllTasksRequest) (*pb.GetAllTasksResponse, error) {
	tasks, err := h.svc.GetAllTasks()
	if err != nil {
		return nil, err
	}

	var pbTasks []*pb.Task
	for _, task := range tasks {
		pbTasks = append(pbTasks, &pb.Task{
			Id:     task.ID,
			Task:   task.Task,
			UserId: task.UserID,
		})
	}

	return &pb.GetAllTasksResponse{Tasks: pbTasks}, nil
}

func (h *Handler) GetTaskByID(ctx context.Context, req *pb.GetTaskByIDRequest) (*pb.GetTaskByIDResponse, error) {
	task, err := h.svc.GetTaskByID(req.Id)
	if err != nil {
		return nil, err
	}

	return &pb.GetTaskByIDResponse{
		Task: &pb.Task{
			Id:     task.ID,
			Task:   task.Task,
			UserId: task.UserID,
		},
	}, nil
}

func (h *Handler) GetTasksByUserID(ctx context.Context, req *pb.GetTasksByUserIDRequest) (*pb.GetTasksByUserIDResponse, error) {
	tasks, err := h.svc.GetTasksByUserID(req.UserId)
	if err != nil {
		return nil, err
	}

	var pbTasks []*pb.Task
	for _, task := range tasks {
		pbTasks = append(pbTasks, &pb.Task{
			Id:     task.ID,
			Task:   task.Task,
			UserId: task.UserID,
		})
	}

	return &pb.GetTasksByUserIDResponse{Tasks: pbTasks}, nil
}

func (h *Handler) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskResponse, error) {
	task, err := h.svc.UpdateTask(req.Id, req.Task)
	if err != nil {
		return nil, err
	}

	return &pb.UpdateTaskResponse{
		Task: &pb.Task{
			Id:     task.ID,
			Task:   task.Task,
			UserId: task.UserID,
		},
	}, nil
}

func (h *Handler) DeleteTask(ctx context.Context, req *pb.DeleteTaskRequest) (*emptypb.Empty, error) {
	err := h.svc.DeleteTask(req.Id)
	if err != nil {
		return nil, err
	}
	return &emptypb.Empty{}, nil
}
