package routes

import (
	"github.com/Aanu1995/restaurant-management/controllers"
	"github.com/Aanu1995/restaurant-management/middlewares"
	"github.com/gin-gonic/gin"
)

func TableRouter(router *gin.RouterGroup){
	// Food routes
	tableRouter := router.Group("/tables", middlewares.Authenticate)

	tableRouter.GET("", controllers.GetTables)
	tableRouter.GET("/:tableId", controllers.GetTable)
	tableRouter.POST("", controllers.CreateTable)
	tableRouter.PATCH("/:tableId", controllers.UpdateTable)
}