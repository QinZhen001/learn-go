package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// r := gin.Default() // use Logger() and Recovery() middleware
	r := gin.New() // no middleware
	r.GET("/", func(ctx *gin.Context) {
		panic("test panic")
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
		})
	})
	r.Run(":8080")
}
