package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	fb "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/firebase"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/repository"
	"github.com/rs/xid"
)

func RegisterHandler(firebaseApp *fb.FirebaseApp, db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input fb.RegisterUserBody
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "faild to bind json"})
			return
		}
		userID := xid.New().String()
		req := &fb.RegisterUser{ID: userID, Email: input.Email, Password: input.Password, Name: input.Name}
		// firebaseにユーザ登録
		record, err := fb.CreateUserWithUID(ctx, firebaseApp.Client, req)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("ユーザ作成に失敗しました")})
			return
		}

		if err := repository.InsertUser(req, db); err != nil {

		}
		ctx.JSON(http.StatusCreated, gin.H{"message": record.UID})
	}
}
