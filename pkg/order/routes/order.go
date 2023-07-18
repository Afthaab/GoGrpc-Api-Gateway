package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/apiGateway/pkg/domain"
	"github.com/apiGateway/pkg/order/pb"
	"github.com/apiGateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func CreateOrder(ctx *gin.Context, c pb.OrderManagementClient) {
	id, _ := strconv.Atoi(ctx.GetString("userId"))
	orderData := domain.Orders{}
	err := ctx.Bind(&orderData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.CreateOrder(context.Background(), &pb.CreateOrderRequest{
		Userid:      int64(id),
		Paymentid:   orderData.Paymentid,
		Addressid:   orderData.Addid,
		Orderstatus: "Order Placed",
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusUnprocessableEntity, false, "Create order failed", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Create order Successful",
			"data":    res,
		})
	}
}

func SaveCartData(ctx *gin.Context, c pb.OrderManagementClient) {
	id, _ := strconv.Atoi(ctx.GetString("userId"))
	orderItems := domain.OrderItems{}
	err := ctx.BindJSON(&orderItems)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	var datas []*pb.CartItems
	for _, item := range orderItems.Cartproduct {
		data := pb.CartItems{
			Uid:         int64(id),
			Pid:         item.Pid,
			Productname: item.Productname,
			Sizename:    item.Sizename,
			Price:       item.Price,
			Imagelink:   item.Imageurls,
			Quantity:    item.Quantity,
		}
		datas = append(datas, &data)
	}
	res, err := c.SaveCartItems(context.Background(), &pb.CartItemsSaveRequest{
		Orderid: orderItems.Orderid.String(),
		Items:   datas,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusBadRequest, false, "Saving cart product failed", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Saving cart product failed successful",
			"data":    res,
		})
	}
}

func OrderDispatch(ctx *gin.Context, c pb.OrderManagementClient) {
	orderid := ctx.Query("id")
	res, err := c.OrderDispacth(context.Background(), &pb.OrderDispacthRequest{
		Orderid: orderid,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusBadRequest, false, "Update the order details failed", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Update the order details successfull",
			"data":    res,
		})
	}

}

func CancelOrder(ctx *gin.Context, c pb.OrderManagementClient) {
	orderid := ctx.Query("id")
	res, err := c.OrderCancel(context.Background(), &pb.OrderCancelRequest{
		Orderid: orderid,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusBadRequest, false, "Cancel order failed", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Cancel order successfull",
			"data":    res,
		})
	}
}

func GetAllOrders(ctx *gin.Context, c pb.OrderManagementClient) {
	res, err := c.AdminViewAllOrders(context.Background(), &pb.AdminViewAllOrdersRequest{})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusNotFound, false, "Get all orders failed", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Get all orders successfull",
			"data":    res,
		})
	}
}

func ViewOrderById(ctx *gin.Context, c pb.OrderManagementClient) {
	orderid := ctx.Query("id")
	res, err := c.ViewOrderById(context.Background(), &pb.ViewOrderByIdRequest{
		Id: orderid,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusNotFound, false, "Get order by id failed", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Get order by id successfull",
			"data":    res,
		})
	}
}
