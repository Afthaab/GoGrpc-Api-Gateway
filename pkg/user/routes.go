package profile

import (
	"github.com/apiGateway/pkg/auth"
	"github.com/apiGateway/pkg/config"
	"github.com/apiGateway/pkg/user/routes"
	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(r *gin.Engine, cfg *config.Config, authSvc *auth.ServiceAuth) {
	svc := &UserService{
		client: InitUserService(cfg),
	}
	authorize := auth.InitAuthMiddleware(authSvc)

	user := r.Group("/user")

	user.GET("/profile", authorize.AuthRequired, svc.ViewProfile)
	user.PUT("/profile/edit", authorize.AuthRequired, svc.EditProfile)
	user.PATCH("/profile/change/password", authorize.AuthRequired, svc.ChangePassword)

	user.POST("/address/add", authorize.AuthRequired, svc.AddAdress)
	user.GET("/address/view", authorize.AuthRequired, svc.ViewAddress)
	user.PUT("/address/edit", authorize.AuthRequired, svc.EditAddress)

}

func (svc *UserService) ViewProfile(ctx *gin.Context) {
	routes.ViewProfile(ctx, svc.client)
}

func (svc *UserService) EditProfile(ctx *gin.Context) {
	routes.EditProfile(ctx, svc.client)
}
func (svc *UserService) ChangePassword(ctx *gin.Context) {
	routes.ChangePassword(ctx, svc.client)
}

func (svc *UserService) AddAdress(ctx *gin.Context) {
	routes.AddAdress(ctx, svc.client)
}

func (svc *UserService) ViewAddress(ctx *gin.Context) {
	routes.ViewAddress(ctx, svc.client)
}

func (svc *UserService) EditAddress(ctx *gin.Context) {
	routes.EditAddress(ctx, svc.client)
}
