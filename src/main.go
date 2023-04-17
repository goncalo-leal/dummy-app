package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("ping", ping)

	router.Run("0.0.0.0:8080")
}

func ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
