package main

import (
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()

	engine.GET("/sandeep", func(c *gin.Context) {
		c.String(http.StatusOK, "hello")
	})
	



	// run server on PORT
	engine.Run(port())
}

func port() string {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = "8080"
	}
	return ":" + port
}