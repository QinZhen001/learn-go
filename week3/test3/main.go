package main

import (
	"learngo/week3/test3/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func accountList(c *gin.Context) {
	var accountList []model.Account
	al := model.Account{No: 1, Name: "111"}
	a2 := model.Account{No: 2, Name: "222"}
	accountList = append(accountList, al)
	accountList = append(accountList, a2)
	c.JSON(http.StatusOK, gin.H{
		"accountList": accountList,
	})
}

func addAccount(c *gin.Context) {}

func main() {
	r := gin.Default()
	accountGroup := r.Group("/account")
	{
		accountGroup.GET("/list", accountList)
		accountGroup.POST("/add", addAccount)
	}
}
