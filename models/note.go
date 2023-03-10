package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Note struct{
	ID 							primitive.ObjectID		`bson:"_id"`
	NoteId					string								`json:"noteId" bson:"noteId"`
	Title						string								`json:"title"`
	Description			string								`json:"description"`
	CreatedAt 			string								`json:"createdAt" bson:"createdAt"`
	UpdatedAt 			string								`json:"updatedAt" bson:"updatedAt"`
}