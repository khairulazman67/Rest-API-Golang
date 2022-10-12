package controllers

import (
	"errors"
	"net/http"

	"assignment_02/database"
	"assignment_02/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	Orders := models.Order{
		Customer_name: body.Customer_name,
		// Ordered_at:    body.Ordered_at,
	}

	if result := db.Create(&Orders); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"orderedAt":    &Orders.Ordered_at,
		"CustomerName": &Orders.Customer_name,
		"Items":        &Orders,
	},
	)
}

func GetOrder(ctx *gin.Context) {
	db := database.GetDB()

	Order := models.Order{}

	err := db.Find(&Order).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			ctx.AbortWithError(http.StatusNotFound, err)
			return
		}
		ctx.AbortWithError(http.StatusNotFound, err)
		// print("Error finding user : ", err)
	}
	// fmt.Printf("User Data : %+v \n", Order)
	ctx.JSON(http.StatusCreated, gin.H{
		"order": Order,
	},
	)
}

func DeleteOrder(ctx *gin.Context) {
	db := database.GetDB()

	id := ctx.Param("ID")

	order := models.Order{}

	err := db.Where("id =?", id).Delete(&order).Error

	if err != nil {
		ctx.AbortWithError(http.StatusNotFound, err)
		return
	}
	ctx.JSON(http.StatusCreated, gin.H{
		"Status": "Delete data berhasil",
	},
	)
}
