package routes

import (
	"context"
	"net/http"

	"github.com/apiGateway/pkg/auth/pb"
	"github.com/apiGateway/pkg/domain"
	"github.com/apiGateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, p pb.AuthServiceClient) {
	body := domain.User{}

	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error in Binding the JSON Data",
		})
		return
	}

	res, err := p.Register(context.Background(), &pb.RegisterRequest{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})
	if err != nil {
		responses := utils.ErrorResponse("Failed to create user", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*ctx, responses)
		return
	}
	responses := utils.SuccessResponse(true, "SUCCESS", res)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusBadRequest)
	utils.ResponseJSON(*ctx, responses)
	return

}

func RegisterValidate(ctx *gin.Context, p pb.AuthServiceClient) {
	body := domain.User{}
	if err := ctx.Bind(&body); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Error": "Error in Binding the JSON Data",
		})
		return
	}
	res, err := p.RegisterValidate(context.Background(), &pb.RegisterValidateRequest{
		Otp: body.Otp,
	})
	errs, _ := utils.ExtractError(err)
	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Error": errs,
		})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"Message": "Successfully created the user",
		"body":    res,
	})

}
