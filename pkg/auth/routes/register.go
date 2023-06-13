package routes

import (
	"context"
	"fmt"
	"net/http"

	"github.com/apiGateway/pkg/auth/pb"
	"github.com/apiGateway/pkg/domain"
	"github.com/apiGateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func Register(ctx *gin.Context, p pb.AuthServiceClient) {
	body := domain.User{}
	err := ctx.BindJSON(&body)
	fmt.Println(err)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	fmt.Print(body)

	res, err := p.Register(context.Background(), &pb.RegisterRequest{
		Username: body.Username,
		Email:    body.Email,
		Password: body.Password,
	})

	// extracting the error message from the GRPC error
	errs, _ := utils.ExtractError(err)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "Registering the user fialed",
			"Error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Successfully sent the Otp to the user",
			"data":    res,
		})
	}

}

func RegisterValidate(ctx *gin.Context, p pb.AuthServiceClient) {
	body := domain.User{}
	err := ctx.Bind(&body)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := p.RegisterValidate(context.Background(), &pb.RegisterValidateRequest{
		Otp: body.Otp,
	})

	// extracting the error message from the GRPC error
	errs, _ := utils.ExtractError(err)

	if err != nil {
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "OTP Verification Failed",
			"Error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Successfully registered the user",
			"data":    res,
		})
	}

}
