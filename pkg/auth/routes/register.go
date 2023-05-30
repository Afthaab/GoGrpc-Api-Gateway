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
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}
	res, err := p.Regitser(context.Background(), &pb.RegisterRequest{
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
