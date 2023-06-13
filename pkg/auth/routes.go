package auth

import (
	"net/http"

	"github.com/apiGateway/pkg/auth/routes"
	"github.com/apiGateway/pkg/config"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, cfg *config.Config) *ServiceAuth {
	svc := &ServiceAuth{
		client: InitServiceClient(cfg),
	}
	auth := InitAuthMiddleware(svc)
	route := r.Group("/auth")
	route.POST("/register", svc.Register)
	route.POST("/register/validate", svc.RegitserValidate)
	route.POST("/login", svc.Login)
	route.GET("/home", auth.AuthRequired, func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"Message": "Success",
		})
	})
	return svc

}

func (svc *ServiceAuth) Register(ctx *gin.Context) {
	routes.Register(ctx, svc.client)
}

func (svc *ServiceAuth) RegitserValidate(ctx *gin.Context) {
	routes.RegisterValidate(ctx, svc.client)
}

func (svc *ServiceAuth) Login(ctx *gin.Context) {
	routes.Login(ctx, svc.client)
}
