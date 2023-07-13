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
		sizes.PUT("/edit", authorize.AuthRequired, svc.EditSize)
		sizes.DELETE("/delete", authorize.AuthRequired, svc.DeleteSize)
	}
	category := r.Group("/category")
	{
		category.POST("/add", authorize.AuthRequired, svc.AddCategory)
		category.GET("/view", authorize.AuthRequired, svc.ViewCategories)
		category.GET("/view/id", authorize.AuthRequired, svc.ViewCategoryById)
		category.PUT("/edit", authorize.AuthRequired, svc.EditCategory)
		category.DELETE("/delete", authorize.AuthRequired, svc.DeleteCategory)
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
		foodtype.PUT("/edit", authorize.AuthRequired, svc.EditFoodType)
		foodtype.DELETE("/delete", authorize.AuthRequired, svc.DeleteFoodType)
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

func (svc *ProdcutService) ViewCategoryById(ctx *gin.Context) {
	routes.ViewCategoryById(ctx, svc.client)
}

func (svc *ProdcutService) EditCategory(ctx *gin.Context) {
	routes.EditCategory(ctx, svc.client)
}

func (svc *ProdcutService) EditFoodType(ctx *gin.Context) {
	routes.EditFoodType(ctx, svc.client)
}

func (svc *ProdcutService) EditSize(ctx *gin.Context) {
	routes.EditSize(ctx, svc.client)
}

func (svc *ProdcutService) DeleteSize(ctx *gin.Context) {
	routes.DeleteSize(ctx, svc.client)
}

func (svc *ProdcutService) DeleteCategory(ctx *gin.Context) {
	routes.DeleteCategory(ctx, svc.client)
}

func (svc *ProdcutService) DeleteFoodType(ctx *gin.Context) {
	routes.DeleteFoodType(ctx, svc.client)
}
