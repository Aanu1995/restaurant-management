package models

import "go.mongodb.org/mongo-driver/bson/primitive"


type User struct {
	ID 							primitive.ObjectID  `bson:"_id"`
	UserId 					string							`json:"userId" bson:"userId"`
	FirstName 			*string							`json:"firstName" bson:"firstName" validate:"required,min=2,max=100"`
	LastName 				*string							`json:"lastName" bson:"lastName" validate:"required,min=2,max=100"`
	Email 					*string							`json:"email" validate:"required,email"`
	Password				*string							`json:"password" validate:"required,min=6"`
	Phone 					*string							`json:"phone" validate:"required"`
	Avatar					*string							`json:"avatar"`
	AccessToken 		*string							`json:"accessToken" bson:"accessToken"`
	RefreshToken 		*string							`json:"refreshToken" bson:"refreshToken"`
	CreatedAt 			string							`json:"createdAt" bson:"createdAt"`
	UpdatedAt 			string							`json:"updatedAt" bson:"updatedAt"`
}


type Token struct {
	AccessToken 		string							`json:"accessToken" bson:"accessToken"`
	RefreshToken 		string							`json:"refreshToken" bson:"refreshToken"`
}