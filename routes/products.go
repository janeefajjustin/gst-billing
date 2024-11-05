package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/gst-billing/models"
)

func addProduct(context *gin.Context) {

	var product models.Product
	err := context.ShouldBindJSON(&product)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse request data."})
		return
	}

	err = product.Save()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not add the product. Try again later."})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "Product Added!", "Product": product})

}

func getProduct(context *gin.Context) {
	productName := context.Param("prod")
	product, err := models.FetchProductByName(productName)
	if err != nil {
		productCode, err := strconv.ParseInt(productName, 10, 64)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse product code."})
			return
		}
		product, err := models.FetchProductByCode(productCode)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch product."})
			return
		}
		context.JSON(http.StatusOK, product)
		return

	}
	context.JSON(http.StatusOK, product)
}
