package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Invoice struct{
	ID 							primitive.ObjectID		`bson:"_id"`
	InvoiceId				string								`json:"invoiceId" bson:"invoiceId"`
	OrderId					string								`json:"orderId" bson:"orderId"`
	PaymentMethod		*string								`json:"paymentMethod" bson:"paymentMethod" validate:"eq=CARD|eq=CASH|eq="`
	PaymentStatus		*string								`json:"paymentStatus" bson:"paymentStatus" validate:"required,eq=PENDING|eq=PAID"`
	PaymentDueDate	string								`json:"paymentDueDate" bson:"paymentDueDate"`
	CreatedAt 			string								`json:"createdAt" bson:"createdAt"`
	UpdatedAt 			string								`json:"updatedAt" bson:"updatedAt"`
}