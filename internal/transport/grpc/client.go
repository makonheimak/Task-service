package grpc

import (
	pbUser "github.com/makonheimak/project-protos/proto/user"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// NewUserClient создает клиента для UserService (точно по заданию)
func NewUserClient(addr string) (pbUser.UserServiceClient, *grpc.ClientConn, error) {
	// 1. grpc.Dial с insecure credentials (т.к. задание говорит WithInsecure, но в новых версиях WithTransportCredentials)
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, nil, err
	}

	// 2. userpb.NewUserServiceClient(conn)
	client := pbUser.NewUserServiceClient(conn)

	// 3. вернуть client, conn, err
	return client, conn, nil
}
