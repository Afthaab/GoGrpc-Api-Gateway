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
	{
		profile := user.Group("/profile")
		{
			profile.GET("/view", authorize.AuthRequired, svc.ViewProfile)
			profile.PUT("/edit", authorize.AuthRequired, svc.EditProfile)
			profile.PATCH("/change/password", authorize.AuthRequired, svc.ChangePassword)
		}
		address := user.Group("/address")
		{
			address.POST("/add", authorize.AuthRequired, svc.AddAdress)
			address.GET("/view", authorize.AuthRequired, svc.ViewAddressById)
			address.GET("/view/all", authorize.AuthRequired, svc.ViewAddress)
			address.PUT("/edit", authorize.AuthRequired, svc.EditAddress)
		}
	}

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

func (svc *UserService) ViewAddressById(ctx *gin.Context) {
	routes.ViewAddressById(ctx, svc.client)
}
