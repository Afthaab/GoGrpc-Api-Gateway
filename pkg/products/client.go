package products

import (
	"log"

	"github.com/apiGateway/pkg/config"
	"github.com/apiGateway/pkg/products/pb"
	"google.golang.org/grpc"
)

type ProdcutService struct {
	client pb.ProductManagementClient
}

func InitProductService(cfg *config.Config) pb.ProductManagementClient {
	grpcConn, err := grpc.Dial(cfg.Productsvcurl, grpc.WithInsecure())
	if err != nil {
		log.Fatalln("Could not connect to the GRPC Server")
	}
	return pb.NewProductManagementClient(grpcConn)
}
