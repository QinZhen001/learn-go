package main

import (
	"fmt"
	"learngo/week3/test6/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("login", loginHandler)
	r.Run()
}

func loginHandler(c *gin.Context) {
	var login model.Accountlogin
	err := c.ShouldBind(&login)
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusNotFound, gin.H{
			"msg": "输入错误",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"msg": fmt.Sprintf("%s-%s", login.AccountName, login.Password),
	})
}
