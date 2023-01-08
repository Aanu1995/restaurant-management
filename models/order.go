package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Order struct{
	ID 							primitive.ObjectID		`bson:"_id"`
	OrderId					string								`json:"orderId" bson:"orderId"`
	TableId					*string								`json:"tableId" bson:"tableId" validate:"required"`
	OrderDate 			string								`json:"orderDate" bson:"orderDate" validate:"required"`
	CreatedAt 			string								`json:"createdAt" bson:"createdAt"`
	UpdatedAt 			string								`json:"updatedAt" bson:"updatedAt"`
}