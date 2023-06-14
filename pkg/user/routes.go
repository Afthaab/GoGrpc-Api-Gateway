package profile

import (
	"github.com/apiGateway/pkg/auth"
	"github.com/apiGateway/pkg/config"
	"github.com/apiGateway/pkg/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterProileRoutes(r *gin.Engine, cfg *config.Config, authSvc *auth.ServiceAuth) {
	svc := &ProfileService{
		client: InitProfileService(cfg),
	}
	authorize := auth.InitAuthMiddleware(authSvc)

	user := r.Group("/user")

	user.GET("/profile", authorize.AuthRequired, svc.ViewProfile)
}

func (svc *ProfileService) ViewProfile(ctx *gin.Context) {
	routes.ViewProfile(ctx, svc.client)
}
