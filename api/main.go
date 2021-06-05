package main

import (
	"goLoginBoiler/database"
	"goLoginBoiler/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r = handlers.InitHandlers(r)

	database.InitDB()

	r.Run()
}
