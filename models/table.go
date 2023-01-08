package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Table struct{
	ID 							primitive.ObjectID		`bson:"_id"`
	TableId					string								`json:"TableId" bson:"TableId"`
	NumberOfGuests	*int									`json:"numberOfGuests" bson:"numberOfGuests" validate:"required"`
	TableNumber			*int									`json:"tableNumber" bson:"tableNumber" validate:"required"`
	CreatedAt 			string								`json:"createdAt" bson:"createdAt"`
	UpdatedAt 			string								`json:"updatedAt" bson:"updatedAt"`
}