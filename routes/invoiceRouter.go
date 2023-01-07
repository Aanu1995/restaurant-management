package routes

import (
	"github.com/Aanu1995/restaurant-management/middlewares"
	"github.com/gin-gonic/gin"
)

func InvoiceRouter(router *gin.RouterGroup){
	// Food routes
	invoiceRouter := router.Group("/invoices", middlewares.Authenticate)

	invoiceRouter.GET("", controllers.GetInvoices)
	invoiceRouter.GET("/:invoiceId", controllers.GetInvoice)
	invoiceRouter.POST("", controllers.CreateInvoice)
	invoiceRouter.PATCH("/:invoiceId", controllers.UpdateInvoice)
}