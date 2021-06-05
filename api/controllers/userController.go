package controllers

import (
	"goLoginBoiler/auth"
	"goLoginBoiler/database"
	"goLoginBoiler/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Login(context *gin.Context) {

	var loginData models.User
	context.Bind(&loginData)

	var loginPair models.User

	database.GetDB().Where("email = ? ", loginData.Email).First(&loginPair)

	if loginPair.Email == "" {
		context.JSON(401, gin.H{
			"type":    "error",
			"message": "User does not exist",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(loginPair.Password), []byte(loginData.Password)); err != nil {
		context.JSON(401, gin.H{
			"type":    "error",
			"message": "Wrong password",
		})
		return
	}

	expTime := time.Now().Add(time.Hour * 2)

	jwt := auth.JWTWrapper{
		SecretKey:   "weyo",
		Issuer:      "AuthService",
		Expirations: int(expTime.Unix()),
	}

	signedToken, err := jwt.GenerateJWT(loginData.Email)
	if err != nil {
		context.JSON(500, gin.H{
			"type":    "error",
			"message": "There was an error generating your login token",
		})
	}

	http.SetCookie(context.Writer, &http.Cookie{
		Name:     "token",
		Value:    signedToken,
		Expires:  expTime,
		HttpOnly: true,
	})

	context.JSON(200, gin.H{
		"type":    "success",
		"message": "Token generated in Cookie",
	})
}

func Signup(context *gin.Context) {
	var userContainer models.User

	context.Bind(&userContainer)

	if userContainer.Email == "" || userContainer.FirstName == "" ||
		userContainer.LastName == "" || userContainer.Password == "" {
		context.JSON(422, gin.H{
			"type": "error", "messsage": "empty fields in body"})
	}

	var checkUser models.User

	database.GetDB().Where("email = ?", userContainer.Email).First(&checkUser)

	if checkUser.Email != "" {
		context.JSON(409, gin.H{
			"type":    "error",
			"message": "Mail already in use",
		})
		return
	}

	hashedPass, err := bcrypt.GenerateFromPassword([]byte(userContainer.Password), bcrypt.DefaultCost)

	if err != nil {
		context.JSON(500, gin.H{
			"type":    "error",
			"message": "There was an error hashing passwords",
		})

		log.Fatalln("There was an error hashing passwords")
	}

	userContainer.Password = string(hashedPass)

	database.GetDB().Create(&userContainer)

	context.JSON(200, gin.H{
		"type":    "success",
		"message": "Signup for user successful",
	})
}
