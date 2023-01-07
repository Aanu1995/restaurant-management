package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Table struct{
	ID 							primitive.ObjectID		`bson:"_id"`
	TableId					string								`json:"TableId"`
	NumberOfGuests	uint									`json:"numberOfGuests" validate:"required"`
	TableNumber			uint									`json:"tableNumber" validate:"required"`
	CreatedAt 			string								`json:"createdAt"`
	UpdatedAt 			string								`json:"updatedAt"`
}