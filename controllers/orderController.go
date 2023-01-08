package controllers

import (
	"net/http"

	"github.com/Aanu1995/restaurant-management/models"
	"github.com/Aanu1995/restaurant-management/services"
	"github.com/gin-gonic/gin"
)


func GetOrder(ctx *gin.Context){
	orderId := ctx.Param("orderId")

	order, err := services.GetOrder(orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": order})
}

func GetOrders(ctx *gin.Context){
	orders, err := services.GetOrders()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": orders})
}


func CreateOrder(ctx *gin.Context){
	var order models.Order

	if err := ctx.BindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the struct
	if err := validate.Struct(order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder, err := services.CreateOrder(order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newOrder})
}


func UpdateOrder(ctx *gin.Context){
	orderId := ctx.Param("orderId")

	var order models.Order
	if err := ctx.BindJSON(&order); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrder, err := services.UpdateOrder(orderId, order)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newOrder})
}