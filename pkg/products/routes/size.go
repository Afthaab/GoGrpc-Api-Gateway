package routes

import (
	"context"
	"net/http"

	"github.com/apiGateway/pkg/domain"
	"github.com/apiGateway/pkg/products/pb"
	"github.com/apiGateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AddSize(ctx *gin.Context, c pb.ProductManagementClient) {
	sizeData := domain.Size{}
	err := ctx.Bind(&sizeData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.AddSizeBazedPrize(context.Background(), &pb.AddSizeBazedPrizeRequest{
		Name:       sizeData.Name,
		Pricerange: sizeData.Price,
	})
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Adding new size failed",
			"err":     err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Adding new size Successfull",
			"data":    res,
		})
	}
}
