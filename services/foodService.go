package services

import (
	"context"
	"errors"
	"time"

	"github.com/Aanu1995/restaurant-management/database"
	"github.com/Aanu1995/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var foodCollection = database.OpenCollection("food")


func GetFood(foodId string) (food models.Food, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	result := foodCollection.FindOne(ctx, bson.M{"foodId": foodId})
	err = result.Decode(&food)

	return
}

func CreateFood(requestBody models.Food) (food models.Food, err error){
	food = requestBody

	// check if the menu exists
	if isMenuExists := checkIfMenuExists(requestBody.MenuId); !isMenuExists {
		err = errors.New("Menu not found")
		return
	}

	createdAt := time.Now().UTC().Format(time.RFC3339)

	food.CreatedAt = createdAt
	food.UpdatedAt = createdAt
	food.ID = primitive.NewObjectID()
	food.FoodId = food.ID.Hex()

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err = foodCollection.InsertOne(ctx, food)

	return
}