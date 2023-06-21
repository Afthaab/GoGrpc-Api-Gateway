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

func AddAdress(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))

	addressData := domain.Address{}
	err := ctx.Bind(&addressData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.AddAddress(context.Background(), &pb.AddAddressRequest{
		Id:       int64(id),
		Houseno:  addressData.Houseno,
		Area:     addressData.Area,
		Landmark: addressData.Landmark,
		City:     addressData.City,
		Type:     addressData.Type,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Add Address Failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Add Address Successfull",
			"data":    res,
		})
	}

}
