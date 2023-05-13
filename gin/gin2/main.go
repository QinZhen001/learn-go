package main

import (
	"learngo/gin/gin2/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	accountGroup := r.Group("/account")
	{
		accountGroup.GET("/list", accountListHandler)
		accountGroup.POST("/add", addAccountHandler)
	}

	productGroup := r.Group("/product")
	{
		productGroup.GET("/list", productListHandler)
	}

	// ...

	r.Run(":8080")
}

func accountListHandler(ctx *gin.Context) {
	var accounts []model.Account
	a1 := model.Account{
		No:   1,
		Name: "test1",
	}
	accounts = append(accounts, a1)
	a2 := model.Account{
		No:   2,
		Name: "test2",
	}
	accounts = append(accounts, a2)
	ctx.JSON(http.StatusOK, gin.H{
		"accounts": accounts,
	})
}

func addAccountHandler(ctx *gin.Context) {

}

func productListHandler(ctx *gin.Context) {}
