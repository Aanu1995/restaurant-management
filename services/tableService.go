package services

import (
	"context"
	"time"

	"github.com/Aanu1995/restaurant-management/database"
	"github.com/Aanu1995/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
)

var tableCollection = database.OpenCollection("table")

/// --------------------------------------------------------------------------------
/// Helper Functions
/// --------------------------------------------------------------------------------

func checkIfTableExists(tableId string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var table models.Table
	// check if user with phone number exists
	err := tableCollection.FindOne(ctx, bson.M{"tableId": tableId}).Decode(&table);

	return err == nil // if no error then table exists
}