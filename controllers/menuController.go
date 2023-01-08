package controllers

import (
	"net/http"

	"github.com/Aanu1995/restaurant-management/models"
	"github.com/Aanu1995/restaurant-management/services"
	"github.com/gin-gonic/gin"
)


func GetMenu(ctx *gin.Context){
	foodId := ctx.Param("menuId")

	menu, err := services.GetMenu(foodId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": menu})
}

func GetMenus(ctx *gin.Context){
	// get user with userId
	menus, err := services.GetMenus()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": menus})
}


func CreateMenu(ctx *gin.Context){
	var menu models.Menu

	if err := ctx.BindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the struct
	if err := validate.Struct(menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMenu, err := services.CreateMenu(menu)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newMenu})
}


func UpdateMenu(ctx *gin.Context){
	menuId := ctx.Param("menuId")

	var menu models.Menu
	if err := ctx.BindJSON(&menu); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newMenu, err := services.UpdateMenu(menuId, menu)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newMenu})
}