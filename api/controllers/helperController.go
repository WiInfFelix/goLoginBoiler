package controllers

import (
	"fmt"
	"goLoginBoiler/auth"
	"log"

	"github.com/gin-gonic/gin"
)

func Ping(context *gin.Context) {
	context.JSON(200, gin.H{
		"message": "pong",
	})
}

func LoginPing(context *gin.Context) {
	token, err := context.Cookie("token")
	if err != nil {
		log.Println(err)
	}

	claims, err := auth.ValidateToken(token)

	if err != nil {
		context.JSON(401, gin.H{
			"type":    "error",
			"message": "There was an error with your claims",
		})
		return
	}

	username := claims.Email

	context.JSON(200, gin.H{
		"message": fmt.Sprintf("Welcome user %s", username),
	})
}
