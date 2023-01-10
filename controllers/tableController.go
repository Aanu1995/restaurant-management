package controllers

import (
	"net/http"

	"github.com/Aanu1995/restaurant-management/models"
	"github.com/Aanu1995/restaurant-management/services"
	"github.com/gin-gonic/gin"
)

func GetTable(ctx *gin.Context){
	tableId := ctx.Param("tableId")

	table, err := services.GetTable(tableId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": table})
}

func GetTables(ctx *gin.Context){
	tables, err := services.GetTables()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": tables})
}


func CreateTable(ctx *gin.Context){
	var table models.Table

	if err := ctx.BindJSON(&table); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the struct
	if err := validate.Struct(table); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTable, err := services.CreateTable(table)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newTable})
}


func UpdateTable(ctx *gin.Context){
	tableId := ctx.Param("tableId")

	var table models.Table
	if err := ctx.BindJSON(&table); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newTable, err := services.UpdateTable(tableId, table)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newTable})
}