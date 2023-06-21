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
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
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

func EditProfile(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))

	userData := domain.User{}
	err := ctx.Bind(&userData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.EditProfile(context.Background(), &pb.EditProfileRequest{
		Id:       int64(id),
		Username: userData.Username,
		Gender:   userData.Gender,
		Dob:      userData.Dateofbirth,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Edit Profile Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Edit Profile Succeded",
			"data":    res,
		})
	}
}

func ChangePassword(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))

	passwordData := domain.Password{}
	err := ctx.Bind(&passwordData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.ChangePassword(context.Background(), &pb.ChangeRequest{
		Id:          int64(id),
		Oldpassword: passwordData.Oldpassword,
		Newpassword: passwordData.Newpassword,
	})

	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(int(res.Status), gin.H{
			"Success": false,
			"Message": "Edit Profile Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(int(res.Status), gin.H{
			"Success": true,
			"Message": "Edit Profile Failed",
			"data":    res,
		})
	}

}
