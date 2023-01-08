package controllers

import (
	"net/http"
	"strconv"

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
	recordPerPage, err := strconv.Atoi(ctx.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 20
	}

	page, err1 := strconv.Atoi(ctx.Query("page"))
	if err1 != nil || page < 1 {
		page = 1
	}

	foods, err2 := services.GetFoods(recordPerPage, page)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err2.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": foods, "nextPage": page + 1})
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

	ctx.JSON(http.StatusCreated, gin.H{"data": newFood})
}


func UpdateFood(ctx *gin.Context){
	foodId := ctx.Param("foodId")

	var food models.Food
	if err := ctx.BindJSON(&food); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newFood, err := services.UpdateFood(foodId, food)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newFood})
}