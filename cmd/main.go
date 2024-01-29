package main

import (
	"go-sms-verify/api"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// inizialize config
	app := api.Config{Router: router}

	// routes
	app.Routes()

	router.Run(":8080")
}
