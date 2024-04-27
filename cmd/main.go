package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"MTC conn": "ok",
		})
	})

	if err := r.Run(); err != nil {
		panic(fmt.Sprintf("Error while running the server: %v", err.Error()))
	}
}
