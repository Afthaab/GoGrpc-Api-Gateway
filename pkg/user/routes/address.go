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
		Id:              int64(id),
		Type:            addressData.Type,
		Locationaddress: addressData.Locationaddress,
		Completeaddress: addressData.Completeaddress,
		Landmark:        addressData.Landmark,
		Floorno:         addressData.Floorno,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Add address failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Add address successfull",
			"data":    res,
		})
	}

}

func ViewAddress(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))

	res, err := c.ViewAddress(context.Background(), &pb.ViewAddressRequest{
		Id: int64(id),
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "View address failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View address successfull",
			"data":    res,
		})
	}
}

func EditAddress(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))
	addid, _ := strconv.Atoi(ctx.Query("addressid"))
	addressData := domain.Address{}
	err := ctx.Bind(&addressData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.EditAddress(context.Background(), &pb.EditAddressRequest{
		Id:              int64(id),
		Addressid:       int64(addid),
		Type:            addressData.Type,
		Locationaddress: addressData.Locationaddress,
		Completeaddress: addressData.Completeaddress,
		Landmark:        addressData.Landmark,
		Floorno:         addressData.Floorno,
	})

	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Edit address failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Edit address successful",
			"data":    res,
		})
	}
}

func ViewAddressById(ctx *gin.Context, c pb.ProfileManagementClient) {
	// get the id from bearer token
	id, _ := strconv.Atoi(ctx.GetString("userId"))
	addid, _ := strconv.Atoi(ctx.Query("addressid"))
	addressData := domain.Address{}
	err := ctx.Bind(&addressData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}

	res, err := c.ViewAddressById(context.Background(), &pb.ViewAddressByIdRequest{
		Addid: int64(addid),
		Uid:   int64(id),
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)

		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "View address failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View address successfull",
			"data":    res,
		})
	}
}
