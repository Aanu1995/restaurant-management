package routes

import (
	"github.com/Aanu1995/restaurant-management/middlewares"
	"github.com/gin-gonic/gin"
)

func OrderRouter(router *gin.RouterGroup){
	// Food routes
	orderRouter := router.Group("/orders", middlewares.Authenticate)

	orderRouter.GET("", controllers.GetOrders)
	orderRouter.GET("/:orderId", controllers.GetOrder)
	orderRouter.POST("", controllers.CreateOrder)
	orderRouter.PATCH("/:orderId", controllers.UpdateOrder)
}