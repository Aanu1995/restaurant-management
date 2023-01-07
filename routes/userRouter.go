package routes

import (
	"github.com/Aanu1995/restaurant-management/controllers"
	"github.com/Aanu1995/restaurant-management/middlewares"
	"github.com/gin-gonic/gin"
)


func UserRouter(router *gin.RouterGroup){
	// User routes
	userRouter := router.Group("/users", middlewares.Authenticate)

	userRouter.GET("", controllers.GetUsers)
	userRouter.GET("/:userId", controllers.GetUser)
}