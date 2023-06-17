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
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))

	userData := domain.User{}
	err := ctx.Bind(&userData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.ViewProfile(context.Background(), &pb.ViewProfileRequest{
		Id: int64(id),
	})

	// extracting the error message from the GRPC error
	errs, _ := utils.ExtractError(err)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "View Profile Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View Profile Successfull",
			"data":    res,
		})
	}
}
