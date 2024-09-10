package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/iqbal13/hactiv8tugasdua/controllers"
)

func SetupRouter() *gin.Engine {
	app := gin.Default()

	// Order routes
	app.GET("/orders", controllers.GetOrders)
	app.GET("/orders/:id", controllers.GetOrderByID)
	app.POST("/orders", controllers.CreateOrder)
	app.PUT("/orders/:id", controllers.UpdateOrder)
	app.DELETE("/orders/:id", controllers.DeleteOrder)

	return app
}
