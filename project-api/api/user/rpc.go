package user

import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	login "test.com/project-api/api/user/user_grpc"
)

var UserClient login.LoginServiceClient

func initGrpcUserClient() {
	conn, err := grpc.Dial(":8881", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	UserClient = login.NewLoginServiceClient(conn)
}
