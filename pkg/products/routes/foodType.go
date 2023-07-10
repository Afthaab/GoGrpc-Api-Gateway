package routes

import (
	"context"
	"net/http"

	"github.com/apiGateway/pkg/domain"
	"github.com/apiGateway/pkg/products/pb"
	"github.com/apiGateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AddFoodType(ctx *gin.Context, c pb.ProductManagementClient) {
	typeData := domain.Foodtype{}
	err := ctx.Bind(&typeData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.AddFoodType(context.Background(), &pb.AddFoodTypeRequest{
		Name:     typeData.Foodtype,
		Imageurl: typeData.Imageurl,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Adding new Food Type failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Adding new Food Type Successfull",
			"data":    res,
		})
	}
}

func ViewFoodType(ctx *gin.Context, c pb.ProductManagementClient) {
	res, err := c.ViewFoodType(context.Background(), &pb.ViewFoodtypeRequest{})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Viewing Food Type failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Viewing Food Type Successfull",
			"data":    res,
		})
	}
}
