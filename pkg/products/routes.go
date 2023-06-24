package products

import (
	"github.com/apiGateway/pkg/auth"
	"github.com/apiGateway/pkg/config"
	"github.com/apiGateway/pkg/products/routes"
	"github.com/gin-gonic/gin"
)

func RegisterProductRoutes(r *gin.Engine, cfg *config.Config, authSvc *auth.ServiceAuth) {
	svc := &ProdcutService{
		client: InitProductService(cfg),
	}
	authorize := auth.InitAuthMiddleware(authSvc)

	r.POST("/sizes/add", authorize.AuthRequired, svc.AddSize)

}

func (svc *ProdcutService) AddSize(ctx *gin.Context) {
	routes.AddSize(ctx, svc.client)
}
