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
	user.PUT("/profile/edit", authorize.AuthRequired, svc.EditProfile)
	user.PATCH("/profile/change/password", authorize.AuthRequired, svc.ChangePassword)

	user.POST("/address/add", authorize.AuthRequired, svc.AddAdress)

}

func (svc *ProfileService) ViewProfile(ctx *gin.Context) {
	routes.ViewProfile(ctx, svc.client)
}

func (svc *ProfileService) EditProfile(ctx *gin.Context) {
	routes.EditProfile(ctx, svc.client)
}
func (svc *ProfileService) ChangePassword(ctx *gin.Context) {
	routes.ChangePassword(ctx, svc.client)
}

func (svc *ProfileService) AddAdress(ctx *gin.Context) {
	routes.AddAdress(ctx, svc.client)
}
