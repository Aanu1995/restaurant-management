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

var orderCollection = database.OpenCollection("order")


func GetOrder(orderId string) (order models.Order, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err = orderCollection.FindOne(ctx, bson.M{"orderId": orderId}).Decode(&order)

	return
}

func GetOrders() (orders []models.Order, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	result, err := orderCollection.Find(ctx, bson.D{})
	defer result.Close(context.Background())

	if err != nil {
		return
	}

	if err = result.All(context.Background(), &orders); err != nil {
		return
	}

	if orders == nil {
		orders = []models.Order{}
	}
	return
}


func CreateOrder(requestBody models.Order) (order models.Order, err error){
	order = requestBody

	// check if the table exists
	if isTableExists := checkIfTableExists(*requestBody.TableId); !isTableExists {
		err = errors.New("Table not found")
		return
	}

	createdAt := time.Now().UTC().Format(time.RFC3339)

	order.CreatedAt = createdAt
	order.UpdatedAt = createdAt
	order.ID = primitive.NewObjectID()
	order.OrderId = order.ID.Hex()

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err = orderCollection.InsertOne(ctx, order)

	return
}


func UpdateOrder(orderId string, requestBody models.Order) (order models.Order, err error){
	// check if table exists if the tableId is to be updated
	if requestBody.TableId != nil {
		if isTableExists := checkIfTableExists(*requestBody.TableId); !isTableExists {
			err = errors.New("Table not found")
			return
		}
	}

	filter := bson.M{"orderId": orderId}
	updateObj := bson.M{
		"tableId": *requestBody.TableId,
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

	if _, err = orderCollection.UpdateOne(ctx, filter, bson.M{"$set": updateObj}); err != nil {
		return
	}

	// return back the updated order
	order, err = GetOrder(orderId)

	return
}

func OrderItemOrderCreator(tableId *string) (orderId string, err error) {
	var order models.Order

	createdAt := time.Now().UTC().Format(time.RFC3339)

	order.OrderDate = createdAt
	order.TableId = tableId
	order.CreatedAt = createdAt
	order.UpdatedAt = createdAt
	order.ID = primitive.NewObjectID()
	order.OrderId = order.ID.Hex()

	// validate the struct
	if err = validate.Struct(order); err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err = orderCollection.InsertOne(ctx, order)

	orderId = order.OrderId
	return
}


/// --------------------------------------------------------------------------------
/// Helper Functions
/// --------------------------------------------------------------------------------
func checkIfOrderExists(orderId string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var order models.Order
	// check if user with phone number exists
	err := orderCollection.FindOne(ctx, bson.M{"orderId": orderId}).Decode(&order);

	return err == nil // if no error then order exists
}
