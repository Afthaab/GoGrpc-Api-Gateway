package order

import (
	"log"

	"github.com/apiGateway/pkg/config"
	"github.com/apiGateway/pkg/order/pb"
	"google.golang.org/grpc"
)

type OrderService struct {
	client pb.OrderManagementClient
}

func InitOrderService(cfg *config.Config) pb.OrderManagementClient {
	grpcConn, err := grpc.Dial(cfg.Ordersvcurl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to the GRPC Server")
	}
	return pb.NewOrderManagementClient(grpcConn)
}
