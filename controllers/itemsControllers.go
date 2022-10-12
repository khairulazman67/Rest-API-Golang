package controllers

import (
	"assignment_02/database"
	"assignment_02/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type ItemsRequestBody struct {
	Item_code   uint   `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	Items       uint   `json:"item"`
	Order_id    uint   `json:"Order_id"`
}
type Items struct {
	Item_code   uint   `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
	Order_id    uint   `json:"Order_id"`
}

type Data struct {
	Ordered_at    time.Time `json:"Ordered_at"`
	Customer_name string    `json:"Customer_name"`
	Item          []Items
}

func CreateOrdersItems(ctx *gin.Context) {
	db := database.GetDB()

	body := Data{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	var Items models.Item

	for _, v := range body.Item {
		Items = models.Item{
			Item_code:   v.Item_code,
			Description: v.Description,
			Quantity:    v.Quantity,
			Order_id:    v.Order_id,
		}
	}

	Orders := models.Order{
		Customer_name: body.Customer_name,
		Ordered_at:    body.Ordered_at,
	}

	if result := db.Create(&Orders); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	if result := db.Create(&Items); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"Data": body,
	},
	)
}
