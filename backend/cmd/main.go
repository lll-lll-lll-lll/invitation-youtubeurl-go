package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	container "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/container"
)

func main() {
	input := container.Input{ID: "id_1", Password: "password_1", URL: "https://www.youtube.com/watch?v=Ceia3jquZN8"}
	_, err := container.New(input)
	if err != nil {
		log.Fatal(err)
	}
	router := gin.Default()
	router.POST("/", func(ctx *gin.Context) {
		var input container.Input
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "faild to bind json"})
			return
		}
	})
	router.Run(":8080")

	// fmt.Println(string(container.EncryptedText))
	// d := aes.Decrypt(container.Cipher, []byte(container.IV), input.String(), container.EncryptedText)
	// fmt.Println(string(d))
}
