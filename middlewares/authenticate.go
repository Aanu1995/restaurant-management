package middlewares

import (
	"net/http"

	"github.com/Aanu1995/restaurant-management/helpers"
	"github.com/gin-gonic/gin"
)

func Authenticate(ctx *gin.Context){
	clientToken := ctx.GetHeader("token")
	if clientToken == "" {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Authorization header required"})
		ctx.Abort()
		return
	}

	claims, err := helpers.ValidateToken(clientToken)

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		ctx.Abort()
		return
	}

	ctx.Set("userId", claims.UserId)
	ctx.Next()
}