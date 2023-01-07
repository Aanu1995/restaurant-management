package services

import (
	"context"
	"time"

	"github.com/Aanu1995/restaurant-management/database"
	"github.com/Aanu1995/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
)


var menuCollection = database.OpenCollection("menu")


func checkIfMenuExists(menuId string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var menu models.Menu
	// check if user with phone number exists
	err := menuCollection.FindOne(ctx, bson.M{"menuId": menuId}).Decode(&menu);

	return err == nil // if no error then menu exists
}