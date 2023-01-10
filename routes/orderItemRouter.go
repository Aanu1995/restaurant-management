package routes

import (
	"github.com/Aanu1995/restaurant-management/controllers"
	"github.com/Aanu1995/restaurant-management/middlewares"
	"github.com/gin-gonic/gin"
)

func OrderItemRouter(router *gin.RouterGroup){
	// Food routes
	orderItemRouter := router.Group("/orderItems", middlewares.Authenticate)

	orderItemRouter.GET("", controllers.GetOrderItems)
	orderItemRouter.GET("/:orderItemId", controllers.GetOrderItem)
	orderItemRouter.GET("/order/:orderId", controllers.GetOrderItemsByOrderId)
	orderItemRouter.POST("", controllers.CreateOrderItem)
	orderItemRouter.PATCH("/:orderItemId", controllers.UpdateOrderItem)
}