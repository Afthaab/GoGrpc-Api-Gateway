package routes

import (
	"context"
	"fmt"
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

		utils.FailureJson(ctx, http.StatusBadRequest, false, "Adding new Food Type failed", errs)
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

		utils.FailureJson(ctx, http.StatusBadRequest, false, "Viewing Food Type failed", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Viewing Food Type Successfull",
			"data":    res,
		})
	}
}

func EditFoodType(ctx *gin.Context, c pb.ProductManagementClient) {
	typeData := domain.Foodtype{}
	err := ctx.Bind(&typeData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.EditFoodType(context.Background(), &pb.EditFoodTypeRequest{
		Typeid:   typeData.Id,
		Typename: typeData.Foodtype,
		Imageurl: typeData.Imageurl,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusBadRequest, false, "Edit type failed", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Edit type successfull",
			"data":    res,
		})
	}
}

func DeleteFoodType(ctx *gin.Context, c pb.ProductManagementClient) {
	typeid := ctx.Query("id")
	fmt.Println(typeid, "PPPPPPPPPPPPPPPPPPPPPPP")
	res, err := c.DeleteFoodType(context.Background(), &pb.DeleteFoodTypeRequest{
		Typeid: typeid,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusUnprocessableEntity, false, "Could not delete the food type", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Delete food type successfull",
			"data":    res,
		})
	}
}
