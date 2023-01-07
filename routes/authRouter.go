package routes

import (
	"github.com/Aanu1995/restaurant-management/controllers"
	"github.com/gin-gonic/gin"
)


func AuthRouter(router *gin.RouterGroup){
	// Authentication routes
	authRouter := router.Group("/auth")

	authRouter.POST("/signup", controllers.SignUp)
	authRouter.POST("/login", controllers.Login)
}