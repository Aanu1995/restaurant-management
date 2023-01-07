package routes

import (
	"github.com/Aanu1995/restaurant-management/controllers"
	"github.com/Aanu1995/restaurant-management/middlewares"
	"github.com/gin-gonic/gin"
)

func FoodRouter(router *gin.RouterGroup){
	// Food routes
	foodRouter := router.Group("/foods", middlewares.Authenticate)

	foodRouter.GET("", controllers.GetFoods)
	foodRouter.GET("/:foodId", controllers.GetFood)
	foodRouter.POST("", controllers.CreateFood)
	foodRouter.PATCH("/:foodId", controllers.UpdateFood)
}