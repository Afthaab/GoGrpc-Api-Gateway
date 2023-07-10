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

	sizes := r.Group("/size")
	{
		sizes.POST("/add", authorize.AuthRequired, svc.AddSize)
		sizes.GET("/view", authorize.AuthRequired, svc.ViewSize)
	}
	category := r.Group("/category")
	{
		category.POST("/add", authorize.AuthRequired, svc.AddCategory)
		category.GET("/view", authorize.AuthRequired, svc.ViewCategories)
	}
	product := r.Group("/product")
	{
		product.POST("/add", authorize.AuthRequired, svc.AddProducts)
		product.GET("/view", authorize.AuthRequired, svc.ViewProducts)
		product.GET("/view/id", authorize.AuthRequired, svc.ViewProductById)
	}
	foodtype := r.Group("/food/type")
	{
		foodtype.POST("/add", authorize.AuthRequired, svc.AddFoodType)
		foodtype.GET("/view", authorize.AuthRequired, svc.ViewFoodType)
	}
	image := r.Group("/image")
	{
		image.POST("/add", authorize.AuthRequired, svc.AddImage)
	}
}

func (svc *ProdcutService) AddSize(ctx *gin.Context) {
	routes.AddSize(ctx, svc.client)
}

func (svc *ProdcutService) AddCategory(ctx *gin.Context) {
	routes.AddCategory(ctx, svc.client)
}

func (svc *ProdcutService) AddProducts(ctx *gin.Context) {
	routes.AddProducts(ctx, svc.client)
}

func (svc *ProdcutService) ViewSize(ctx *gin.Context) {
	routes.ViewSize(ctx, svc.client)
}

func (svc *ProdcutService) ViewCategories(ctx *gin.Context) {
	routes.ViewCategories(ctx, svc.client)
}

func (svc *ProdcutService) ViewProducts(ctx *gin.Context) {
	routes.ViewProducts(ctx, svc.client)
}

func (svc *ProdcutService) AddImage(ctx *gin.Context) {
	routes.AddImage(ctx, svc.client)
}

func (svc *ProdcutService) AddFoodType(ctx *gin.Context) {
	routes.AddFoodType(ctx, svc.client)
}

func (svc *ProdcutService) ViewFoodType(ctx *gin.Context) {
	routes.ViewFoodType(ctx, svc.client)
}

func (svc *ProdcutService) ViewProductById(ctx *gin.Context) {
	routes.ViewProductById(ctx, svc.client)
}
