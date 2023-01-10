package services

import (
	"context"
	"time"

	"github.com/Aanu1995/restaurant-management/database"
	"github.com/Aanu1995/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var tableCollection = database.OpenCollection("table")


func GetTable(tableId string) (table models.Table, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err = tableCollection.FindOne(ctx, bson.M{"tableId": tableId}).Decode(&table)

	return
}

func GetTables() (tables []models.Table, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	result, err := tableCollection.Find(ctx, bson.D{})
	defer result.Close(context.Background())

	if err != nil {
		return
	}

	if err = result.All(context.Background(), &tables); err != nil {
		return
	}

	if tables == nil {
		tables = []models.Table{}
	}
	return
}


func CreateTable(requestBody models.Table) (table models.Table, err error){
	table = requestBody

	createdAt := time.Now().UTC().Format(time.RFC3339)

	table.CreatedAt = createdAt
	table.UpdatedAt = createdAt
	table.ID = primitive.NewObjectID()
	table.TableId = table.ID.Hex()

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err = tableCollection.InsertOne(ctx, table)

	return
}


func UpdateTable(tableId string, requestBody models.Table) (table models.Table, err error){

	filter := bson.M{"tableId": tableId}
	updateObj := bson.M{
		"numberOfGuests": *requestBody.NumberOfGuests,
		"tableNumber": *requestBody.TableNumber,
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

	if _, err = tableCollection.UpdateOne(ctx, filter, bson.M{"$set": updateObj}); err != nil {
		return
	}

	// return back the updated table
	table, err = GetTable(tableId)

	return
}


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