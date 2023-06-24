package profile

import (
	"log"

	"github.com/apiGateway/pkg/config"
	"github.com/apiGateway/pkg/user/pb"
	"google.golang.org/grpc"
)

type UserService struct {
	client pb.ProfileManagementClient
}

func InitUserService(cfg *config.Config) pb.ProfileManagementClient {
	grpcConn, err := grpc.Dial(cfg.Usersvcurl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to the GRPC Server")
	}
	return pb.NewProfileManagementClient(grpcConn)
}
