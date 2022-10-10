package controllers

import (
	"net/http"

	"assignment_02/database"
	"assignment_02/models"

	"github.com/gin-gonic/gin"
)

type ItemsRequestBody struct {
	Item_code   uint   `json:"item_code"`
	Description string `json:"description"`
	Quantity    uint   `json:"quantity"`
}

func CreateItems(ctx *gin.Context) {
	db := database.GetDB()
	body := ItemsRequestBody{}

	if err := ctx.BindJSON(&body); err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

	Items := models.Items{
		Item_code:   body.Item_code,
		Description: body.Description,
		Quantity:    body.Quantity,
	}

	if result := db.Create(&Items); result.Error != nil {
		ctx.AbortWithError(http.StatusNotFound, result.Error)
		return
	}

	ctx.JSON(http.StatusCreated, &Items)
}
