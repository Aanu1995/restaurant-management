package helpers

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type SignedDetails struct {
	UserId				string
	jwt.RegisteredClaims
}

var accessTokenKey = os.Getenv("ACCESS_TOKEN_KEY")
var refreshTokenKey = os.Getenv("REFRESH_TOKEN_KEY")

func GenerateTokens(userId string) (accessToken string, refreshToken string, err error){
	claims := SignedDetails{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
	}

	refreshClaims := SignedDetails{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(168 * time.Hour)),
		},
	}

	if accessToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte(accessTokenKey)); err != nil {
		return
	}

	refreshToken, err = jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString([]byte(refreshTokenKey))
	return
}


func ValidateToken(signedToken string) (claims *SignedDetails, err error){
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(accessTokenKey), nil
	})

	if err != nil {
		return
	}


	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		err  = errors.New("Invalid authorization token")
		return
	}

	if claims.ExpiresAt.Before(time.Now()){
		err  = errors.New("Token is Expired")
		return
	}

	return
}

func ValidateRefreshToken(signedToken string) (claims *SignedDetails, err error){
	token, err := jwt.ParseWithClaims(signedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(refreshTokenKey), nil
	})

	if err != nil {
		return
	}


	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		err  = errors.New("Invalid authorization token")
		return
	}

	if claims.ExpiresAt.Before(time.Now()){
		err  = errors.New("Token is Expired")
		return
	}

	return
}