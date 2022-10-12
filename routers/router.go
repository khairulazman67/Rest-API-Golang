package routers

import (
	"assignment_02/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/orders", controllers.CreateOrdersItems)
	// router.POST("/orders", controllers.CreateOrders)
	router.GET("/orders", controllers.GetOrder)

	router.DELETE("/orders/:ID", controllers.DeleteOrder)
	return router
}
