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
	user := r.Group("/user")
	{
		user.POST("/register", svc.Register)
		user.POST("/register/validate", svc.RegitserValidate)
		user.POST("/login", svc.Login)
		forget := user.Group("forget")
		{
			forget.POST("/forget/password", svc.ForgotPassword)
			forget.POST("/forget/password/validate", svc.RegitserValidate)
			forget.POST("/forget/password/validate/newpassword", svc.ChangePassword)
		}
		user.POST("/forget/password", svc.ForgotPassword)
		user.POST("/forget/password/validate", svc.RegitserValidate)
		user.POST("/forget/password/validate/newpassword", svc.ChangePassword)
	}
	admin := r.Group("/admin")
	{
		admin.POST("/login", svc.AdminLogin)
		admin.POST("/forget/password", svc.ForgotPassword)
		admin.POST("/forget/password/validate", svc.RegitserValidate)
		admin.POST("/forget/password/validate/newpassword", svc.ChangePassword)
	}

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

func (svc *ServiceAuth) AdminLogin(ctx *gin.Context) {
	routes.AdminLogin(ctx, svc.client)
}

func (svc *ServiceAuth) ForgotPassword(ctx *gin.Context) {
	routes.ForgotPassword(ctx, svc.client)
}

func (svc *ServiceAuth) ChangePassword(ctx *gin.Context) {
	routes.ChangePassword(ctx, svc.client)
}
