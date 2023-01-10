package services

import (
	"context"
	"time"

	"github.com/Aanu1995/restaurant-management/database"
	"github.com/Aanu1995/restaurant-management/models"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var orderItemCollection = database.OpenCollection("orderItem")
var validate = validator.New()


func GetOrderItem(orderItemId string) (orderItem models.OrderItem, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err = orderItemCollection.FindOne(ctx, bson.M{"orderItemId": orderItemId}).Decode(&orderItem)

	return
}

func GetOrderItems() (orderItems []models.OrderItem, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	result, err := orderItemCollection.Find(ctx, bson.D{})
	defer result.Close(context.Background())

	if err != nil {
		return
	}

	if err = result.All(context.Background(), &orderItems); err != nil {
		return
	}

	if orderItems == nil {
		orderItems = []models.OrderItem{}
	}
	return
}

func GetOrderItemsByOrderId(orderId string) (orderItems []primitive.M, err error) {
	matchStage := bson.D{{Key: "$match", Value: bson.D{{Key: "orderId", Value: orderId}}}}
	lookupStage := bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "food"},
			{Key: "localField", Value: "foodId"},
			{Key: "foreignField", Value: "foodId"},
			{Key: "as", Value: "food"},
		},
	}}
	unwindStage := bson.D{{Key: "$unwind", Value: bson.D{
		{Key: "path", Value: "$food"},
		{Key: "preserveNullAndEmptyArrays", Value: true},
	}}}

	lookupOrderStage := bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "order"},
			{Key: "localField", Value: "orderId"},
			{Key: "foreignField", Value: "orderId"},
			{Key: "as", Value: "order"},
		},
	}}
	unwindOrderStage := bson.D{{Key: "$unwind", Value: bson.D{
		{Key: "path", Value: "$order"},
		{Key: "preserveNullAndEmptyArrays", Value: true},
	}}}

	lookupTableStage := bson.D{{Key: "$lookup", Value: bson.D{
			{Key: "from", Value: "table"},
			{Key: "localField", Value: "order.tableId"},
			{Key: "foreignField", Value: "tableId"},
			{Key: "as", Value: "table"},
		},
	}}
	unwindTableStage := bson.D{{Key: "$unwind", Value: bson.D{
		{Key: "path", Value: "$table"},
		{Key: "preserveNullAndEmptyArrays", Value: true},
	}}}

	projectStage := bson.D{{Key: "$project", Value: bson.D{
		{Key: "_id", Value: 0},
		{Key: "amount", Value: "$food.price"},
		{Key: "quantity", Value: 1},
		{Key: "foodName", Value: "$food.name"},
		{Key: "price", Value: "$food.price"},
		{Key: "foodImage", Value: "$food.foodImage"},
		{Key: "tableId", Value: "$table.tableId"},
		{Key: "orderId", Value: "$order.orderId"},
		{Key: "tableNumber", Value: "$table.tableNumber"},
		{Key: "numberOfGuests", Value: "$table.numberOfGuests"},
	}}}

	groupStage := bson.D{{Key: "$group", Value: bson.D{
		{Key: "_id", Value: bson.D{
			{Key: "orderId", Value: "$orderId"},
			{Key: "tableId", Value: "$tableId"},
			{Key: "tableNumber", Value: "$tableNumber"},
		}},
		{Key: "paymentDue", Value: bson.D{{Key: "$sum", Value: "$amount"}}},
		{Key: "totalCount", Value: bson.D{{Key: "$sum", Value: 1}}},
		{Key: "orderItems", Value: bson.D{{Key: "$push", Value: "$$ROOT"}}},
	}}}

	projectStage2 := bson.D{{Key: "$project", Value: bson.D{
		{Key: "_id", Value: 0},
		{Key: "paymentDue", Value: 1},
		{Key: "totalCount", Value: 1},
		{Key: "tableNumber", Value: "$_id.tableNumber"},
		{Key: "orderItems", Value: 1},
	}}}

	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	result, err := orderItemCollection.Aggregate(ctx, mongo.Pipeline{
		matchStage,
		lookupStage,
		unwindStage,
		lookupOrderStage,
		unwindOrderStage,
		lookupTableStage,
		unwindTableStage,
		projectStage,
		groupStage,
		projectStage2,
	})
	defer result.Close(context.Background())

	if err != nil {
		return
	}

	err = result.All(context.Background(), &orderItems)

	return
}


func CreateOrderItem(requestBody models.OrderItemPack) (orderItemPack models.OrderItemPack, err error){
	orderItemPack = requestBody

	orderId, err := OrderItemOrderCreator(orderItemPack.TableId)
	if err != nil {
		return
	}

	orderItems := []models.OrderItem{}

	for _, orderItem := range orderItemPack.OrderItems {
		createdAt := time.Now().UTC().Format(time.RFC3339)

		orderItem.CreatedAt = createdAt
		orderItem.UpdatedAt = createdAt
		orderItem.ID = primitive.NewObjectID()
		orderItem.OrderItemId = orderItem.ID.Hex()
		orderItem.OrderId = orderId

		// validate the struct
		if err = validate.Struct(orderItem); err != nil {
			return
		}

		orderItems = append(orderItems, orderItem)
	}

	// cast type of models.OrderItem to interface{}
	var interfaceSlice []interface{} = make([]interface{}, len(orderItems))
	for index, item := range orderItems {
		interfaceSlice[index] = item
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()
	_, err = orderItemCollection.InsertMany(ctx, interfaceSlice)

	if err != nil{
		return
	}

	orderItemPack.OrderItems = orderItems

	return
}


func UpdateOrderItem(orderItemId string, requestBody models.OrderItem) (orderItem models.OrderItem, err error){

	filter := bson.M{"orderItemId": orderItemId}
	updateObj := bson.M{
		"quantity": *requestBody.Quantity,
		"price": *requestBody.Price,
		"foodId": *requestBody.FoodId,
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

	if _, err = orderItemCollection.UpdateOne(ctx, filter, bson.M{"$set": updateObj}); err != nil {
		return
	}

	// return back the updated orderItem
	orderItem, err = GetOrderItem(orderItemId)

	return
}