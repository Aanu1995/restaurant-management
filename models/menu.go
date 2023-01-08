package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Menu struct{
	ID 							primitive.ObjectID		`bson:"_id"`
	MenuId					string								`json:"menuId" bson:"menuId"`
	Name						*string								`json:"name" validate:"required"`
	Category				*string								`json:"category" validate:"required"`
	StartDate 			*string								`json:"startDate" bson:"startDate"`
	EndDate 				*string								`json:"endDate" bson:"endDate"`
	CreatedAt 			string								`json:"createdAt" bson:"createdAt"`
	UpdatedAt 			string								`json:"updatedAt" bson:"updatedAt"`
}