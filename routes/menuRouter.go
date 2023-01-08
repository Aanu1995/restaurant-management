package routes

import (
	"github.com/Aanu1995/restaurant-management/controllers"
	"github.com/Aanu1995/restaurant-management/middlewares"
	"github.com/gin-gonic/gin"
)

func MenuRouter(router *gin.RouterGroup){
	// Food routes
	menuRouter := router.Group("/menus", middlewares.Authenticate)

	menuRouter.GET("", controllers.GetMenus)
	menuRouter.GET("/:menuId", controllers.GetMenu)
	menuRouter.POST("", controllers.CreateMenu)
	menuRouter.PATCH("/:menuId", controllers.UpdateMenu)
}