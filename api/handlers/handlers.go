package handlers

import (
	"goLoginBoiler/controllers"

	"github.com/gin-gonic/gin"
)

func InitHandlers(router *gin.Engine) *gin.Engine {

	router.POST("/login", controllers.Login)
	router.POST("/signup", controllers.Signup)
	router.POST("/loginPing", controllers.LoginPing)
	router.GET("/ping", controllers.Ping)

	return router
}
