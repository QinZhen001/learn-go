package main

import (
	"fmt"
	"learngo/gin/gin4/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/login", loginHandler)
	r.Run(":8080")
}

func loginHandler(ctx *gin.Context) {
	var login model.Account
	err := ctx.ShouldBind(&login)
	if err != nil {
		fmt.Println(err)
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "error login",
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"msg":   "ok",
		"login": login,
	})
}
