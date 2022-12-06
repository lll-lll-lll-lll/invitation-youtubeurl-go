package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/handler"
	container "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/container"
	fb "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/firebase"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/repository"

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
	router.POST("/invitation_code", func(ctx *gin.Context) {
		var input container.Input
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "faild to bind json"})
			return
		}
		con, err := container.New(input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		if err := insert(con, postgresql.Db); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": con.Code})
	})
	router.POST("/invitation_users", func(ctx *gin.Context) {
	})
	router.POST("/register", handler.RegisterHandler(firebaseApp, postgresql.Db))
	router.Run(":8080")
}

type InputInvitation struct {
	Code string `json:"code" validate:"required"`
}

func insert(con *container.Container, db *sqlx.DB) error {
	postCode := PostCode{Code: con.Code}
	if err := repository.Transaction(db, postCode, insertCodeFunc); err != nil {
		return err
	}
	return nil
}

type PostCode struct {
	Code string `json:"code"`
}

func insertCodeFunc(req interface{}, db *sqlx.DB) error {
	castedReq := req.(PostCode)
	stmt, err := db.Prepare("INSERT INTO invitation_codes(code) VALUES($1) returning code")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(castedReq.Code)
	if err != nil {
		return err
	}
	return nil
}
