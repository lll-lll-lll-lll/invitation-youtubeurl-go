package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	container "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/container"
	db "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/db"
)

func main() {
	router := gin.Default()
	db := db.NewPostgreSql()
	db.Open()
	router.POST("/", func(ctx *gin.Context) {
		var input container.Input
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "faild to bind json"})
			return
		}
	})
	router.Run(":8080")
}
