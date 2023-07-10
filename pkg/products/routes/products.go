package routes

import (
	"context"
	"net/http"

	"github.com/apiGateway/pkg/domain"
	"github.com/apiGateway/pkg/products/pb"
	"github.com/apiGateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AddProducts(ctx *gin.Context, c pb.ProductManagementClient) {
	productData := domain.Products{}
	err := ctx.Bind(&productData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.AddProduct(context.Background(), &pb.AddProductRequest{
		Name:        productData.Productname,
		Calories:    productData.Calories,
		Availibilty: productData.Availibility,
		Categoryid:  productData.Categoryid,
		Typeid:      productData.Typeid,
		Baseprice:   float32(productData.Baseprice),
		Sizeid:      productData.Sizeid,
		Imagelink:   productData.Imageurls,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Adding new product failed",
			"Error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Adding new product Successfull",
			"Eata":    res,
		})
	}

}

func ViewProducts(ctx *gin.Context, c pb.ProductManagementClient) {
	res, err := c.ViewProduct(context.Background(), &pb.ViewProductRequest{})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "View Product Failed",
			"Error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View Product Successfull",
			"data":    res,
		})
	}
}

func ViewProductById(ctx *gin.Context, c pb.ProductManagementClient) {
	pid := ctx.Query("id")
	res, err := c.ViewProductById(context.Background(), &pb.ViewProductByIdRequest{
		Pid: pid,
	})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusNotFound, gin.H{
			"Success": false,
			"Message": "View product by id failed",
			"Error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View product by id successfull",
			"data":    res,
		})
	}
}
