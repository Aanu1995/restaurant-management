package controllers

import (
	"net/http"

	"github.com/Aanu1995/restaurant-management/models"
	"github.com/Aanu1995/restaurant-management/services"
	"github.com/gin-gonic/gin"
)


func GetOrderItem(ctx *gin.Context){
	orderItemId := ctx.Param("orderItemId")

	orderItem, err := services.GetOrderItem(orderItemId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": orderItem})
}

func GetOrderItems(ctx *gin.Context){
	orderItems, err := services.GetOrderItems()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": orderItems})
}


func GetOrderItemsByOrderId(ctx *gin.Context){
	orderId := ctx.Param("orderId")

	orderItems, err := services.GetOrderItemsByOrderId(orderId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": orderItems})
}

func CreateOrderItem(ctx *gin.Context){
	var orderItemPack models.OrderItemPack

	if err := ctx.BindJSON(&orderItemPack); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the struct
	if err := validate.Struct(orderItemPack); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrderItemPack, err := services.CreateOrderItem(orderItemPack)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newOrderItemPack})
}


func UpdateOrderItem(ctx *gin.Context){
	orderItemId := ctx.Param("orderItemId")

	var orderItem models.OrderItem
	if err := ctx.BindJSON(&orderItem); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newOrderItem, err := services.UpdateOrderItem(orderItemId, orderItem)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newOrderItem})
}
