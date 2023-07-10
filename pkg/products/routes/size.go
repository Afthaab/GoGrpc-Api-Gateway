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
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Adding new size failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Adding new size Successfull",
			"data":    res,
		})
	}
}

func ViewSize(ctx *gin.Context, c pb.ProductManagementClient) {
	res, err := c.ViewSizeBasedPrize(context.Background(), &pb.ViewSizeBasedPriceRequest{})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "View Size Based Prices Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Successfully viewed Size Based Prices",
			"data":    res,
		})
	}
}
