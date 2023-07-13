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
			"Error":   errs,
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
			"Error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Successfully viewed Size Based Prices",
			"data":    res,
		})
	}
}

func EditSize(ctx *gin.Context, c pb.ProductManagementClient) {
	sizeData := domain.Size{}
	err := ctx.Bind(&sizeData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.EditSizeBasedPrize(context.Background(), &pb.EditSizeBasedPrizeRequest{
		Sizeid:     sizeData.ID,
		Sizename:   sizeData.Name,
		Pricerange: sizeData.Price,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusBadRequest, false, "Could not edit the size data", errs)

	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Edit Size Based Price Successfull",
			"data":    res,
		})
	}
}

func DeleteSize(ctx *gin.Context, c pb.ProductManagementClient) {
	sid := ctx.Query("id")
	res, err := c.DeleteSizeBasedPrize(context.Background(), &pb.DeleteSizeBasedPrizeRequest{
		Sizeid: sid,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusUnprocessableEntity, false, "Could not delete the size", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Delete Size Based Price Successfull",
			"data":    res,
		})
	}
}
