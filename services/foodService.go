package services

import (
	"context"
	"errors"
	"time"

	"github.com/Aanu1995/restaurant-management/database"
	"github.com/Aanu1995/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var foodCollection = database.OpenCollection("foods")


func GetFood(foodId string) (food models.Food, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err = foodCollection.FindOne(ctx, bson.M{"foodId": foodId}).Decode(&food)

	return
}

func GetFoods(recordPerPage int, page int) (foods []models.Food, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	opts := options.Find()
	opts.SetSkip(int64((page - 1) * recordPerPage))
	opts.SetLimit(int64(recordPerPage))

	result, err := foodCollection.Find(ctx, bson.D{}, opts)
	defer result.Close(context.Background())

	if err != nil {
		return
	}

	if err = result.All(context.Background(), &foods); err != nil {
		return
	}

	if foods == nil {
		foods = []models.Food{}
	}
	return
}


func CreateFood(requestBody models.Food) (food models.Food, err error){
	food = requestBody

	// check if the menu exists
	if isMenuExists := checkIfMenuExists(*requestBody.MenuId); !isMenuExists {
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


func UpdateFood(foodId string, requestBody models.Food) (food models.Food, err error){
	// check if menu exists if the menuId is to be updated
	if requestBody.MenuId != nil {
		if isMenuExists := checkIfMenuExists(*requestBody.MenuId); !isMenuExists {
			err = errors.New("Menu not found")
			return
		}
	}

	filter := bson.M{"foodId": foodId}
	updateObj := bson.M{
		"name":	*requestBody.Name,
		"price": *requestBody.Price,
		"foodImage": *requestBody.FoodImage,
		"menuId": *requestBody.MenuId,
		"updatedAt": time.Now().UTC().Format(time.RFC3339),
	}

	// delete object is not provided in the request body
	for k, v := range updateObj {
		if v == nil {
			delete(updateObj, k)
		}
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	if _, err = foodCollection.UpdateOne(ctx, filter, bson.M{"$set": updateObj}); err != nil {
		return
	}

	// return back the updated food
	food, err = GetFood(foodId)

	return
}