package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type OrderItem struct{
	ID 							primitive.ObjectID		`bson:"_id"`
	OrderItemId			string								`json:"orderItemId"`
	OrderId					string								`json:"orderId" validate:"required"`
	FoodId					string								`json:"foodId" validate:"required"`
	Quantity				string								`json:"quantity" validate:"required,eq=S|eq=M|eq=L"`
	Price						float64								`json:"price" validate:"required"`
	CreatedAt 			string								`json:"createdAt"`
	UpdatedAt 			string								`json:"updatedAt"`
}