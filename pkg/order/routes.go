package order

import (
	"github.com/apiGateway/pkg/auth"
	"github.com/apiGateway/pkg/config"
	"github.com/apiGateway/pkg/order/routes"
	"github.com/gin-gonic/gin"
)

func RegisterOrderRoutes(r *gin.Engine, cfg *config.Config, authSvc *auth.ServiceAuth) {
	svc := &OrderService{
		client: InitOrderService(cfg),
	}
	authorize := auth.InitAuthMiddleware(authSvc)

	user := r.Group("/user")
	{
		payment := user.Group("/payment")
		{
			payment.POST("/cod", authorize.AuthRequired, svc.CashOnDelivery)
		}
		order := user.Group("/order")
		{
			order.POST("/create", authorize.AuthRequired, svc.CreateOrder)
			order.POST("/save/items", authorize.AuthRequired, svc.SaveCartData)
		}
	}
	admin := r.Group("/admin")
	{
		order := admin.Group("/order")
		order.PATCH("/dispatch", authorize.AuthRequired, svc.OrderDispatch)
		order.PUT("/cancel", authorize.AuthRequired, svc.CancelOrder)
		order.GET("/view/all", authorize.AuthRequired, svc.GetAllOrders)
		order.GET("/view", authorize.AuthRequired, svc.ViewOrderById)
	}

}

func (svc *OrderService) GetAllOrders(ctx *gin.Context) {
	routes.GetAllOrders(ctx, svc.client)
}

func (svc *OrderService) SaveCartData(ctx *gin.Context) {
	routes.SaveCartData(ctx, svc.client)
}

func (svc *OrderService) CashOnDelivery(ctx *gin.Context) {
	routes.CashOnDelivery(ctx, svc.client)
}

func (svc *OrderService) CreateOrder(ctx *gin.Context) {
	routes.CreateOrder(ctx, svc.client)
}

func (svc *OrderService) OrderDispatch(ctx *gin.Context) {
	routes.OrderDispatch(ctx, svc.client)
}

func (svc *OrderService) CancelOrder(ctx *gin.Context) {
	routes.CancelOrder(ctx, svc.client)
}

func (svc *OrderService) ViewOrderById(ctx *gin.Context) {
	routes.ViewOrderById(ctx, svc.client)
}
