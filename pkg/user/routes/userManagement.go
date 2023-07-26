package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apiGateway/pkg/domain"
	"github.com/apiGateway/pkg/user/pb"
	"github.com/apiGateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func ViewAllUsers(ctx *gin.Context, c pb.ProfileManagementClient) {
	res, err := c.ViewAllUsers(context.Background(), &pb.ViewAllUsersRequest{})
	fmt.Println(res, "ppp")
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusBadRequest, false, "View all users failed", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View all users successfull",
			"data":    res,
		})
	}
}

func UserBlockUnblock(ctx *gin.Context, c pb.ProfileManagementClient) {
	userData := domain.User{}
	err := ctx.Bind(&userData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.BlockOrUnblockUser(context.Background(), &pb.BlockRequest{
		Userid:      userData.Id,
		Blockstatus: userData.Isblocked,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		utils.FailureJson(ctx, http.StatusBadRequest, false, "User block/unblock failed", errs)
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "User block/unblock successfull",
			"data":    res,
		})
	}

}
