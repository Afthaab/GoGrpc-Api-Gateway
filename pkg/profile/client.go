package profile

import (
	"log"

	"github.com/apiGateway/pkg/config"
	"github.com/apiGateway/pkg/profile/pb"
	"google.golang.org/grpc"
)

type ProfileService struct {
	client pb.ProfileManagementClient
}

func InitProfileService(cfg *config.Config) pb.ProfileManagementClient {
	grpcConn, err := grpc.Dial(cfg.Profilesvcurl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to the GRPC Server")
	}
	return pb.NewProfileManagementClient(grpcConn)
}
