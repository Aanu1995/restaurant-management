package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type Menu struct{
	ID 							primitive.ObjectID		`bson:"_id"`
	MenuId					string								`json:"menuId" validate:"required"`
	Name						string								`json:"name" validate:"required"`
	Category				string								`json:"category" validate:"required"`
	StartDate 			string								`json:"startDate"`
	EndDate 				string								`json:"endDate"`
	CreatedAt 			string								`json:"createdAt"`
	UpdatedAt 			string								`json:"updatedAt"`
}