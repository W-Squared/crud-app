package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	authorized := router.Group("/", gin.BasicAuth(gin.Accounts{
		"user": "pass",
	}))

	// Serve frontend static files
	authorized.StaticFS("/", http.Dir("./views"))

	router.Run(":8080")
}
