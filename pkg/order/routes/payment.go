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

func CashOnDelivery(ctx *gin.Context, c pb.OrderManagementClient) {
	id, _ := strconv.Atoi(ctx.GetString("userId"))
	paymentData := domain.PaymentData{}
	err := ctx.Bind(&paymentData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.CashOnDelivery(context.Background(), &pb.CashOnDeliveryRequest{
		Totalamount:   paymentData.Totalamount,
		Paymentmode:   "Cash on Delivery",
		Paymentstatus: "Pending",
		Userid:        int64(id),
		Coupondid:     paymentData.Couponid,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusUnprocessableEntity, false, "Cash on delivery failed", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Cash on delivery Successful",
			"data":    res,
		})
	}

}
