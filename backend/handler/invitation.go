package handler

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/container"
	fb "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/firebase"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/repository"
)

// Invitation 招待コードを生成するハンドラー
func Invitation(app *fb.FirebaseApp, db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input container.Input
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "faild to bind json"})
			return
		}
		//ヘッダーからtoken取得
		token, err := fb.GetTokenContext(ctx)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("auth token doesn't exist. %v", err.Error())})
			return
		}
		//firebaseからユーザ情報を取得
		user, err := app.GetUser(ctx, token.UID)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": fmt.Errorf("user not found. %v", err.Error())})
			return
		}
		con, err := container.New(input)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		// 招待コードをinsert
		if err := insertInvitationCode(con, db); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}

		// ユーザを外部キーに持ったinvitationテーブルにインサート
		if err := InsertInvitationCodeWithUser(user.UID, con, db); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": con.Code})
	}
}

type InvitationCodeWithUser struct {
	UserID         string `json:"id"`
	InvitationCode string `json:"invitation_code"`
	IV             string `json:"iv"`
	Key            string `json:"key"`
	EncryptedText  string `json:"encrypted_text"`
}

func InsertInvitationCodeWithUserFunc(req interface{}, db *sqlx.DB) error {
	castedReq := req.(InvitationCodeWithUser)
	stmt, err := db.Prepare("INSERT INTO invitation(id, invitation_code, iv, key,encrypted_text ) VALUES($1,$2,$3,$4,$5)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(castedReq.UserID, castedReq.InvitationCode, castedReq.IV, castedReq.Key, castedReq.EncryptedText)
	if err != nil {
		return err
	}
	return nil
}

func InsertInvitationCodeWithUser(userID string, con *container.Container, db *sqlx.DB) error {
	input := InvitationCodeWithUser{UserID: userID, InvitationCode: con.Code, IV: con.IV.ToHexString(), Key: con.Key, EncryptedText: con.EncryptedText.ToHexString()}
	if err := repository.Transaction(db, input, InsertInvitationCodeWithUserFunc); err != nil {
		return err
	}
	return nil
}

func insertInvitationCode(con *container.Container, db *sqlx.DB) error {
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
