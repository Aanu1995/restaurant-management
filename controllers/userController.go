package controllers

import (
	"net/http"
	"strconv"

	"github.com/Aanu1995/restaurant-management/helpers"
	"github.com/Aanu1995/restaurant-management/models"
	"github.com/Aanu1995/restaurant-management/services"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

var validate = validator.New()

func SignUp(ctx *gin.Context){
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate the struct
	if err := validate.Struct(user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create user account
	if err := services.CreateUser(user); err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"message": "Account created successfully"})
}


func Login(ctx *gin.Context){
	var user models.User

	if err := ctx.BindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// create user account
	newUser, err := services.Login(user)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": newUser})
}


func GetUsers(ctx *gin.Context){
	if err := helpers.CheckUserType(ctx, "ADMIN"); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	recordPerPage, err := strconv.Atoi(ctx.Query("recordPerPage"))
	if err != nil || recordPerPage < 1 {
		recordPerPage = 20
	}

	page, err1 := strconv.Atoi(ctx.Query("page"))
	if err1 != nil || page < 1 {
		page = 1
	}

	// get user with userId
	users, err2 := services.GetUsers(recordPerPage, page)
	if err2 != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": users, "nextPage": page + 1})

}


func GetUser(ctx *gin.Context){
	userId := ctx.Param("userId")

	if err := helpers.MatchUserTypeToUserID(ctx, userId); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error":  err.Error()})
		return
	}

	// get user with userId
	user, err := services.GetUser(userId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}