package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorld(c *gin.Context) {
	c.String(http.StatusOK, "hello world!   dsad")
	dasdasdsa
}

func main() {
	r := gin.Default()

	r.GET("/hello", HelloWorld)

	r.Run(":8080")
}
