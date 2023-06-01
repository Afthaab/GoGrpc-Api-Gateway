package routes

import (
	"context"
	"net/http"

	"github.com/apiGateway/pkg/auth/pb"
	"github.com/apiGateway/pkg/domain"
	"github.com/apiGateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Login(ctx *gin.Context, c pb.AuthServiceClient) {
	user := domain.User{}
	err := ctx.Bind(&user)
	if err != nil {
		responses := utils.ErrorResponse("Please give essential data", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*ctx, responses)
		return
	}
	res, err := c.Login(context.Background(), &pb.LoginRequest{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	})
	if err != nil {
		responses := utils.ErrorResponse("Failed to Login user", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*ctx, responses)
		return
	}
	responses := utils.SuccessResponse(true, "SUCCESS", res)
	ctx.Writer.Header().Set("Content-Type", "application/json")
	ctx.Writer.WriteHeader(http.StatusOK)
	utils.ResponseJSON(*ctx, responses)
	return
}
