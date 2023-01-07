package main

import (
	"net/http"
	"os"

	"github.com/Aanu1995/restaurant-management/routes"
	"github.com/gin-gonic/gin"
)


func main(){
	// get the port number from environment variable
	port := os.Getenv("PORT")

	if port == "" {
		port = "3000"
	}

	router := gin.Default()

	// Group all API routes
	apiRouter := router.Group("/api")

	// Authentication routes
	routes.AuthRouter(apiRouter)
	// User routes
	routes.UserRouter(apiRouter)
	// User routes
	routes.FoodRouter(apiRouter)
	// Invoice routes
	routes.InvoiceRouter(apiRouter)
	// Menu routes
	routes.MenuRouter(apiRouter)
	// Order routes
	routes.OrderRouter(apiRouter)
	// Table routes
	routes.TableRouter(apiRouter)
	// Table routes
	routes.OrderItemRouter(apiRouter)


	apiRouter.GET("", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message":"Welcome to API version 1"})
	})

	router.Run(":" + port)
}