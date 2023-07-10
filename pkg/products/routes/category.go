package routes

import (
	"context"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/apiGateway/pkg/domain"
	"github.com/apiGateway/pkg/products/pb"
	"github.com/apiGateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

func AddImage(ctx *gin.Context, c pb.ProductManagementClient) {
	// Getting the Category Image first
	file, err := ctx.FormFile("image")
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Could not get the image",
			"error":   err,
		})
		return
	}

	// Validate file type
	allowedTypes := map[string]bool{
		".jpeg": true,
		".png":  true,
		".webp": true,
		".jpg":  true,
	}

	// Get the file extension
	ext := strings.ToLower(filepath.Ext(file.Filename))

	// Check if the file type is allowed
	if !allowedTypes[ext] {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Invalid file type",
		})
		return
	}

	// Open the uploaded file
	src, err := file.Open()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Could not open the image",
			"Error":   err.Error(),
		})
		return
	}
	defer src.Close()

	// Read the file as bytes
	fileBytes, err := ioutil.ReadAll(src)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"Success": false,
			"Message": "Could not read the image",
			"Error":   err.Error(),
		})
		return
	}

	// Use the fileBytes for further processing or storing

	res, err := c.AddImage(context.Background(), &pb.AddImageRequest{ImageData: fileBytes})
	if err != nil {
		// Extract the error message from the gRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Adding image failed",
			"error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Adding image success",
			"data":    res,
		})
	}
}

func AddCategory(ctx *gin.Context, c pb.ProductManagementClient) {
	categoryData := domain.Category{}
	err := ctx.Bind(&categoryData)
	if err != nil {
		utils.JsonInputValidation(ctx)
		return
	}
	res, err := c.AddCategories(context.Background(), &pb.AddCategoriesRequest{
		CateggoryName: categoryData.Categoryname,
		Imageurl:      categoryData.Imageurl,
	})

	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "Adding new Category failed",
			"error":   errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "Adding new Category Successfull",
			"data":    res,
		})
	}
}

func ViewCategories(ctx *gin.Context, c pb.ProductManagementClient) {
	res, err := c.ViewCategories(context.Background(), &pb.ViewCategoriesRequest{})
	if err != nil {
		// extracting the error message from the GRPC error
		errs, _ := utils.ExtractError(err)
		ctx.JSON(http.StatusBadRequest, gin.H{
			"Success": false,
			"Message": "View Categories failed",
			"err":     errs,
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"Success": true,
			"Message": "View Categories Successfull",
			"data":    res,
		})
	}
}
