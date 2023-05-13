package main

import (
	"learngo/gin/gin3/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	productGroup := r.Group("/product")
	{
		productGroup.GET("/list", productListHandler)
		productGroup.GET("/:id", productDetailHandler)
		productGroup.GET("/:id/:name", productDetail2Handler)
		productGroup.GET("/file/*all", fileHandler) // /product/file/a/b/c => a/b/c 不常用
	}
	r.Run(":8080")
}

func productListHandler(ctx *gin.Context) {

}

func productDetailHandler(ctx *gin.Context) {
	id := ctx.Param("id")
	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": "id is required",
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"id":  id,
			"msg": "ok",
		})
	}

}

func fileHandler(ctx *gin.Context) {
	all := ctx.Param("all")
	ctx.JSON(http.StatusOK, gin.H{
		"msg": "ok",
		"all": all,
	})
}

func productDetail2Handler(ctx *gin.Context) {
	var p model.Product
	err := ctx.ShouldBindUri(&p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"msg": err.Error(),
		})
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"msg": "ok",
			"p":   p,
		})
	}
}
