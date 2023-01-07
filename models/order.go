package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Order struct{
	ID 							primitive.ObjectID		`bson:"_id"`
	OrderId					string								`json:"orderId"`
	TableId					string								`json:"tableId" validate:"required"`
	OrderDate 			string								`json:"orderDate" validate:"required"`
	CreatedAt 			string								`json:"createdAt"`
	UpdatedAt 			string								`json:"updatedAt"`
}