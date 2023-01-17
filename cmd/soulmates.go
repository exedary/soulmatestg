package main

import (
	"github.com/gin-gonic/gin"
)

func Run() {
}

func main() {
	app := gin.Default()

	//pair.RegisterHandlers(app)

	app.Run(":8080")
}
