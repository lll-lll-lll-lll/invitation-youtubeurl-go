package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/handler"
	fb "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/firebase"
	"log"

	db "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/db"
)

func main() {
	router := gin.Default()
	postgresql := db.NewPostgreSql()
	postgresql.Open()
	defer postgresql.Db.Close()
	firebaseApp, err := fb.InitFireBase()
	if err != nil {
		log.Fatal(err.Error())
	}
	router.POST("/invitation_code", fb.FireBaseAuthRequired(firebaseApp), handler.Invitation(firebaseApp, postgresql.Db))
	router.POST("/register", handler.RegisterHandler(firebaseApp, postgresql.Db))
	router.Run(":8080")
}
