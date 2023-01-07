package services

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/Aanu1995/restaurant-management/database"
	"github.com/Aanu1995/restaurant-management/helpers"
	"github.com/Aanu1995/restaurant-management/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var userCollection *mongo.Collection = database.OpenCollection("Users")



func GetUser(userId string) (user models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	result := userCollection.FindOne(ctx, bson.M{"userid": userId})

	err = result.Decode(&user)

	return
}

func GetUsers(recordPerPage int, page int) (users []models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	opts := options.Find()
	opts.SetSkip(int64((page - 1) * recordPerPage))
	opts.SetLimit(int64(recordPerPage))

	result, err := userCollection.Find(ctx, bson.D{}, opts)
	defer result.Close(context.Background())

	if err != nil {
		return
	}

	if err = result.All(context.Background(), &users); err != nil {
		return
	}

	if users == nil {
		users = []models.User{}
	}
	return
}

func GetUserWithEmail(email string) (user models.User, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	result := userCollection.FindOne(ctx, bson.M{"email": email})

	err = result.Decode(&user)

	return
}

func CreateUser(requestBody models.User) (err error) {
	var user models.User = requestBody

	// check if user with email exists
	if emailExists := userWithEmailExists(user.Email); emailExists {
		err = errors.New("User with this email or phone already exists")
		return
	}

	// check if user with phone exists
	if phoneExists := userWithPhoneExists(user.Phone); phoneExists {
		err = errors.New("User with this email or phone already exists")
		return
	}

	// hash the password
	hashPassword := helpers.Hashpassword(user.Password)
	createdAt := time.Now().UTC().Format(time.RFC3339)

	user.Password = hashPassword
	user.CreatedAt = createdAt
	user.UpdatedAt = createdAt
	user.ID = primitive.NewObjectID()
	user.UserId = user.ID.Hex()

	// get access and refresh token
	accessToken, refreshToken, err := helpers.GenerateTokens(user);
	if err != nil {
		log.Panic("Token generation failure")
		return
	}

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err = userCollection.InsertOne(ctx, user)

	return
}

func Login(requestBody models.User) (user models.User, err error){
	// Get the data of the user with the email provided
	user, err = GetUserWithEmail(requestBody.Email)
	if err != nil {
		err = errors.New("Incorrect email or password")
		return
	}

	// Verify if the user password in database is the same as the password
	// supplied by the user
	passwordIsValid := helpers.VerifyPassword(user.Password, user.Password)
	if !passwordIsValid {
		err = errors.New("Incorrect email or password")
		return
	}

	// Generate and Update the user tokens in database
	err = generateAndUpdateUserTokens(&user);
	if err != nil {
		log.Panic("Token generation failure")
		return
	}

	return
}



/// --------------------------------------------------------------------------------
/// Helper Functions
/// --------------------------------------------------------------------------------

func userWithEmailExists(email string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var user models.User
	// check if user with phone number exists
	err := userCollection.FindOne(ctx, bson.M{"email": email}).Decode(&user);

	return err == nil // if no error then user exists
}

func userWithPhoneExists(phone string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var user models.User
	// check if user with phone number exists
	err := userCollection.FindOne(ctx, bson.M{"phone": phone}).Decode(&user);

	return err == nil // if no error then user exists
}

func generateAndUpdateUserTokens(user *models.User) (err error) {
	accessToken, refreshToken, err := helpers.GenerateTokens(*user);
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	updatedAt := time.Now().UTC().Format(time.RFC3339)
	update := bson.M{"$set": bson.M{"accesstoken": accessToken, "refreshtoken": refreshToken, "updatedat":updatedAt}}
	_, err = userCollection.UpdateOne(ctx, bson.M{"userid": user.UserId}, update)
	if err != nil {
		return
	}

	user.AccessToken = accessToken
	user.RefreshToken = refreshToken
	user.UpdatedAt = updatedAt

	return
}
