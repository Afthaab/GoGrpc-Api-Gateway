package auth

import (
	"github.com/apiGateway/pkg/auth/routes"
	"github.com/apiGateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) *ServiceAuth {
	svc := &ServiceAuth{
		client: InitServiceClient(cfg),
	}
	route := r.Group("/auth")
	route.POST("/register", svc.Register)
	return svc

}

func (svc *ServiceAuth) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.client)
}
