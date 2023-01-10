package services

import (
	"context"
	"errors"
	"time"

	"github.com/Aanu1995/restaurant-management/database"
	"github.com/Aanu1995/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var invoiceCollection = database.OpenCollection("invoice")


func GetInvoice(invoiceId string) (invoiceView models.InvoiceViewFormat, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var invoice models.Invoice
	err = invoiceCollection.FindOne(ctx, bson.M{"invoiceId": invoiceId}).Decode(&invoice)
	if err != nil{
		return
	}

	// Get all the order items of this invoice
	orderItems, err := GetOrderItemsByOrderId(invoice.OrderId)
	if err != nil{
		return
	}

	invoiceView = models.InvoiceViewFormat{
		InvoiceId: invoice.InvoiceId,
		OrderId: invoice.OrderId,
		PaymentDueDate: invoice.PaymentDueDate,
		PaymentStatus: invoice.PaymentStatus,
		TableNumber: orderItems[0]["tableNumber"],
		PaymentDue: orderItems[0]["paymentDue"],
		OrderDetails: orderItems[0]["orderItems"],
	}

	if invoice.PaymentMethod != nil {
			invoiceView.PaymentMethod = *invoice.PaymentMethod
	}

	return
}

func GetInvoices() (invoices []models.Invoice, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	result, err := invoiceCollection.Find(ctx, bson.D{})
	defer result.Close(context.Background())

	if err != nil {
		return
	}

	if err = result.All(context.Background(), &invoices); err != nil {
		return
	}

	if invoices == nil {
		invoices = []models.Invoice{}
	}
	return
}


func CreateInvoice(requestBody models.Invoice) (invoice models.Invoice, err error){
	invoice = requestBody

	// check if the order exists
	if isOrderExists := checkIfOrderExists(requestBody.OrderId); !isOrderExists {
		err = errors.New("Order not found")
		return
	}

	createdAt := time.Now().UTC().Format(time.RFC3339)

	invoice.CreatedAt = createdAt
	invoice.UpdatedAt = createdAt
	invoice.PaymentDueDate = time.Now().AddDate(0, 0, 1).Format(time.RFC3339)
	invoice.ID = primitive.NewObjectID()
	invoice.InvoiceId = invoice.ID.Hex()

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err = invoiceCollection.InsertOne(ctx, invoice)

	return
}


func UpdateInvoice(invoiceId string, requestBody models.Invoice) (invoiceView models.InvoiceViewFormat, err error){
	filter := bson.M{"invoiceId": invoiceId}
	updateObj := bson.M{
		"paymentMethod": *requestBody.PaymentMethod,
		"paymentStatus": *requestBody.PaymentStatus,
		"updatedAt": time.Now().UTC().Format(time.RFC3339),
	}

	// delete object is not provided in the request body
	for k, v := range updateObj {
		if v == nil {
			delete(updateObj, k)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	if _, err = invoiceCollection.UpdateOne(ctx, filter, bson.M{"$set": updateObj}); err != nil {
		return
	}

	// return back the updated invoice
	invoiceView, err = GetInvoice(invoiceId)

	return
}
