package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/handler"
	fb "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/firebase"

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
	router.POST("/get_invitation_code", handler.GetYoutubeURLByInvitationCode(postgresql.Db))
	router.POST("/create_invitation_code", fb.FirebaseMiddleware(firebaseApp), handler.Invitation(firebaseApp, postgresql.Db))
	router.POST("/create_invitation_code_guest", handler.InvitationGuest(postgresql.Db))
	router.POST("/register", handler.Register(firebaseApp, postgresql.Db))
	router.Run(":8080")
}
