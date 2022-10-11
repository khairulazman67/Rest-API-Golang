package controllers

import (
	"net/http"

	"assignment_02/database"
	"assignment_02/models"

	"github.com/gin-gonic/gin"
)

type OrdersRequestBody struct {
	Customer_name string `json:"Customer_name"`
	// Ordered_at    time.Time `json:"quantity"`
}

func CreateOrders(ctx *gin.Context) {
	db := database.GetDB()
	body := OrdersRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	Orders := models.Orders{
		Customer_name: body.Customer_name,
		// Ordered_at:    body.Ordered_at,
	}

	if result := db.Create(&Orders); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &Orders)
}
