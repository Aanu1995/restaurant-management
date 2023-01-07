package controllers

import (
	"net/http"

	"github.com/Aanu1995/restaurant-management/models"
	"github.com/Aanu1995/restaurant-management/services"
	"github.com/gin-gonic/gin"
)


func GetFood(ctx *gin.Context){
	foodId := ctx.Param("foodId")

	food, err := services.GetFood(foodId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": food})
}

func GetFoods(ctx *gin.Context){

}


func CreateFood(ctx *gin.Context){
	var food models.Food

	if err := ctx.BindJSON(&food); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the struct
	if err := validate.Struct(food); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newFood, err := services.CreateFood(food)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newFood})
}


func UpdateFood(ctx *gin.Context){
	
}