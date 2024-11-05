package routes

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/janeefajjustin/gst-billing/models"
)

func generateBill(context *gin.Context) {

	productName := context.Param("prod")
	quantity, err := strconv.ParseInt(context.Param("quantity"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse quantity."})
		return
	}
	product, err := models.FetchProductByName(productName)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not find the product."})
		return
	}

	total, tax, price := calcBilling(product.Price, product.Gst, quantity)
	context.JSON(http.StatusOK,
		gin.H{"Product Name": productName, "Quantity": quantity, "Price of one item": product.Price, "Price (Without tax)": price, "Tax": tax, "Total Price (Includes tax)": total})
	err = product.SaveBilling(quantity, total)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Error in Saving"})
		return
	}
}

func calcBilling(price float64, gst float64, quantity int64) (float64, float64, float64) {
	price = price * float64(quantity)
	tax := (price * gst) / 100
	total := price + tax
	return total, tax, price
}


func getBillings(context *gin.Context) {
	events, err := models.GetAllBilling()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Could not fetch data. Try again later."})
		return
	}
	context.JSON(http.StatusOK, events)
}


