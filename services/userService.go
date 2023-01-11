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

var userCollection *mongo.Collection = database.OpenCollection("users")


func GetUser(userId string) (user models.UserOnly, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	result := userCollection.FindOne(ctx, bson.M{"userId": userId})

	err = result.Decode(&user)

	return
}

func GetUsers(recordPerPage int, page int) (users []models.UserOnly, err error) {
	ctx, cancel := context.WithTimeout(context.Background(), 20 * time.Second)
	defer cancel()

	opts := options.Find()
	opts.SetSkip(int64((page - 1) * recordPerPage))
	opts.SetLimit(int64(recordPerPage))

	result, err := userCollection.Find(ctx, bson.D{}, opts)
	if err != nil {
		return
	}

	defer result.Close(context.Background())

	if err = result.All(context.Background(), &users); err != nil {
		return
	}

	if users == nil {
		users = []models.UserOnly{}
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

	// check if user with email or phone exists
	if userExists := userWithEmailOrPhoneExists(*user.Email, *user.Phone); userExists {
		err = errors.New("user with this email or phone already exists")
		return
	}

	// hash the password
	hashPassword := helpers.Hashpassword(*user.Password)
	createdAt := time.Now().UTC().Format(time.RFC3339)

	user.Password = &hashPassword
	user.CreatedAt = createdAt
	user.UpdatedAt = createdAt
	user.ID = primitive.NewObjectID()
	user.UserId = user.ID.Hex()

	// get access and refresh token
	accessToken, refreshToken, err := helpers.GenerateTokens(user.UserId);
	if err != nil {
		log.Panic("Token generation failure")
		return
	}

	user.AccessToken = &accessToken
	user.RefreshToken = &refreshToken

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	_, err = userCollection.InsertOne(ctx, user)

	return
}

func Login(requestBody models.User) (user models.User, err error){
	// Get the data of the user with the email provided
	user, err = GetUserWithEmail(*requestBody.Email)
	if err != nil {
		err = errors.New("incorrect email or password")
		return
	}

	// Verify if the user password in database is the same as the password
	// supplied by the user
	passwordIsValid := helpers.VerifyPassword(*user.Password, *requestBody.Password)
	if !passwordIsValid {
		err = errors.New("incorrect email or password")
		return
	}

	// Generate and Update the user tokens in database
	err = generateAndUpdateUserTokens(&user);
	if err != nil {
		err = errors.New("token generation failure")
		return
	}

	return
}


func RefreshToken(refreshToken string) (token models.Token, err error){
	claims, err := helpers.ValidateRefreshToken(refreshToken)
	if err != nil {
		return
	}

	user := models.User{
		UserId: claims.UserId,
	}

	// Generate and Update the user tokens in database
	err = generateAndUpdateUserTokens(&user);
	if err != nil {
		log.Panic("Token generation failure")
		return
	}

	token = models.Token{
		AccessToken: *user.AccessToken,
		RefreshToken: *user.RefreshToken,
	}
	return
}


/// --------------------------------------------------------------------------------
/// Helper Functions
/// --------------------------------------------------------------------------------

func userWithEmailOrPhoneExists(email string, phone string) bool {
	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	var user models.User

	filter := bson.D{{
		Key: "$or", Value: bson.A{
			bson.D{{Key: "email", Value: email}},
			bson.D{{Key: "phone", Value: phone}},
		},
	}}
	// check if user with phone number exists
	err := userCollection.FindOne(ctx, filter).Decode(&user);

	return err == nil // if no error then user exists
}

func generateAndUpdateUserTokens(user *models.User) (err error) {
	accessToken, refreshToken, err := helpers.GenerateTokens(user.UserId);
	if err != nil {
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10 * time.Second)
	defer cancel()

	updatedAt := time.Now().UTC().Format(time.RFC3339)
	update := bson.M{"$set": bson.M{"accessToken": accessToken, "refreshToken": refreshToken, "updatedAt":updatedAt}}
	_, err = userCollection.UpdateOne(ctx, bson.M{"userId": user.UserId}, update)
	if err != nil {
		return
	}

	user.AccessToken = &accessToken
	user.RefreshToken = &refreshToken
	user.UpdatedAt = updatedAt

	return
}
