package controllers

import (
	"net/http"

	"github.com/Aanu1995/restaurant-management/models"
	"github.com/Aanu1995/restaurant-management/services"
	"github.com/gin-gonic/gin"
)


func GetInvoice(ctx *gin.Context){
	invoiceId := ctx.Param("invoiceId")

	invoice, err := services.GetInvoice(invoiceId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": invoice})
}

func GetInvoices(ctx *gin.Context){
	invoices, err := services.GetInvoices()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": invoices})
}


func CreateInvoice(ctx *gin.Context){
	var invoice models.Invoice

	if err := ctx.BindJSON(&invoice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the struct
	if err := validate.Struct(invoice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newInvoice, err := services.CreateInvoice(invoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"data": newInvoice})
}


func UpdateInvoice(ctx *gin.Context){
	invoiceId := ctx.Param("invoiceId")

	var invoice models.Invoice
	if err := ctx.BindJSON(&invoice); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	newInvoice, err := services.UpdateInvoice(invoiceId, invoice)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newInvoice})
}