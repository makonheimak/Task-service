package grpc

import (
	"log"
	"net"

	"github.com/makonheimak/task-service/internal/task/service"

	pb "github.com/makonheimak/project-protos/proto/task"
	pbUser "github.com/makonheimak/project-protos/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// RunGRPC запускает gRPC сервер (точно по заданию)
func RunGRPC(svc *service.Service, uc pbUser.UserServiceClient) error {
	// 1. net.Listen на ":50051" (task-service порт)
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		return err
	}

	// 2. grpc.NewServer()
	grpcSrv := grpc.NewServer()

	// 3. taskpb.RegisterTaskServiceServer(grpcSrv, NewHandler(svc, uc))
	handler := NewHandler(svc, uc)
	pb.RegisterTaskServiceServer(grpcSrv, handler)

	// 4. grpcSrv.Serve(listener)
	reflection.Register(grpcSrv)

	log.Println("🚀 Task Service gRPC server starting on :50051")
	return grpcSrv.Serve(lis)
}
