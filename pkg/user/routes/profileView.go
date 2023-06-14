package routes

import (
	"context"
	"net/http"
	"strconv"

	"github.com/apiGateway/pkg/domain"
	"github.com/apiGateway/pkg/user/pb"
	"github.com/apiGateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ViewProfile(ctx *gin.Context, c pb.ProfileManagementClient) {
	id, _ := strconv.Atoi(ctx.GetString("userId"))
	userData := domain.User{}
	err := ctx.Bind(&userData)
	if err != nil {
		responses := utils.ErrorResponse("Please give essential data", err.Error(), nil)
		ctx.Writer.Header().Set("Content-Type", "application/json")
		ctx.Writer.WriteHeader(http.StatusBadRequest)
		utils.ResponseJSON(*ctx, responses)
		return
	}
	res, err := c.ViewProfile(context.Background(), &pb.ViewProfileRequest{
		Id: int64(id),
	})
	if err != nil {
		responses := utils.ErrorResponse("Failed to View the profile", err.Error(), nil)
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
