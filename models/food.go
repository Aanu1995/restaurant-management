package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Food struct{
	ID 							primitive.ObjectID		`bson:"_id"`
	FoodId					string								`json:"foodId" bson:"foodId"`
	MenuId					*string								`json:"menuId" bson:"menuId" validate:"required"`
	Name						*string								`json:"name" validate:"required,min=2"`
	Price						*float64							`json:"price" validate:"required"`
	FoodImage				*string								`json:"foodImage" bson:"foodImage" validate:"required"`
	CreatedAt 			string								`json:"createdAt" bson:"createdAt"`
	UpdatedAt 			string								`json:"updatedAt" bson:"updatedAt"`
}