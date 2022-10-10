package routers

import (
	"assignment_02/controllers"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/items", controllers.CreateItems)
	return router
}
