package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Invoice struct{
	ID 							primitive.ObjectID		`bson:"_id"`
	InvoiceId				string								`json:"invoiceId"`
	OrderId					string								`json:"orderId"`
	PaymentMethod		string								`json:"paymentMethod" validate:"eq=CARD|eq=CASH|eq="`
	PaymentStatus		string								`json:"paymentStatus" validate:"required,eq=PENDING|eq=PAID"`
	PaymentDueDate	string								`json:"paymentDueDate"`
	CreatedAt 			string								`json:"createdAt"`
	UpdatedAt 			string								`json:"updatedAt"`
}