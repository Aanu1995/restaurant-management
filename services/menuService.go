package services

import (
	"context"
	"errors"
	"reflect"
	"time"

	"github.com/Aanu1995/restaurant-management/database"
	"github.com/Aanu1995/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)


var menuCollection = database.OpenCollection("menus")


func GetMenu(menuId string) (menu models.Menu, err error){
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	err = menuCollection.FindOne(ctx, bson.M{"menuId": menuId}).Decode(&menu)

	return
}

func GetMenus() (menus []models.Menu, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()

	result, err := menuCollection.Find(ctx, bson.D{})
	if err != nil {
		return
	}

	defer result.Close(context.Background())

	if err = result.All(context.Background(), &menus); err != nil {
		return
	}

	if menus == nil {
		menus = []models.Menu{}
	}
	return
}


func CreateMenu(requestBody models.Menu) (menu models.Menu, err error){
	menu = requestBody

	createdAt := time.Now().UTC().Format(time.RFC3339)

	menu.CreatedAt = createdAt
	menu.UpdatedAt = createdAt
	menu.ID = primitive.NewObjectID()
	menu.MenuId = menu.ID.Hex()

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err = menuCollection.InsertOne(ctx, menu)

	return
}


func UpdateMenu(menuId string, requestBody models.Menu) (menu models.Menu, err error){
	// check if the correct start date and end date is provided
	if requestBody.StartDate != nil || requestBody.EndDate != nil {
		var startDate, endDate time.Time

		if startDate, err = time.Parse(time.RFC3339, *requestBody.StartDate); err != nil{
			return
		} else if endDate, err = time.Parse(time.RFC3339, *requestBody.EndDate); err != nil{
			return
		} else if !(startDate.After(time.Now()) && endDate.After(startDate)){
			err = errors.New("invalid start date or end date")
			return
		}
	}

	filter := bson.M{"menuId": menuId}
	updateObj := bson.M{
		"name": requestBody.Name,
		"category": requestBody.Category,
		"startDate": requestBody.StartDate,
		"endDate": requestBody.EndDate,
	}

	// delete object is not provided in the request body
	for k, v := range updateObj {
		if reflect.ValueOf(v).IsNil() {
			delete(updateObj, k)
		}
	}

	// update the date
	updateObj["updatedAt"]= time.Now().UTC().Format(time.RFC3339)

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	if _, err = menuCollection.UpdateOne(ctx, filter, bson.M{"$set": updateObj}); err != nil {
		return
	}

	// return back the updated menu
	menu, err = GetMenu(menuId)

	return
}


/// --------------------------------------------------------------------------------
/// Helper Functions
/// --------------------------------------------------------------------------------
func checkIfMenuExists(menuId string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var menu models.Menu
	// check if user with phone number exists
	err := menuCollection.FindOne(ctx, bson.M{"menuId": menuId}).Decode(&menu);

	return err == nil // if no error then menu exists
}