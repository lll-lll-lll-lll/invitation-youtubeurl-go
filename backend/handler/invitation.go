package handler

import (
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/container"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/crypto"
	aes "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/crypto"
	fb "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/firebase"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/repository"
)

type GetInvitationInput struct {
	Code     string `json:"code"`
	ID       string `json:"id"`
	Password string `json:"password"`
}

type InvitationBody struct {
	UserID        string `db:"id" json:"id"`
	Code          string `db:"invitation_code" json:"invitation_code"`
	IV            string `db:"iv" json:"iv"`
	Key           string `db:"key" json:"key"`
	EncryptedText string `db:"encrypted_text" json:"encrypted_text"`
	YoutubeURL    string `db:"url" json:"youtube_url"`
}

func (ib GetInvitationInput) String() string {
	return fmt.Sprintf("%s.%s", ib.ID, ib.Password)
}

func GetYoutubeURLByInvitationCode(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input GetInvitationInput
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "faild to bind json"})
			return
		}
		var invitationBody InvitationBody
		if err := db.QueryRowx("SELECT * FROM invitation_guest WHERE invitation_code = $1", input.Code).StructScan(&invitationBody); err != nil {
			if err == sql.ErrNoRows {
				ctx.JSON(http.StatusBadRequest, fmt.Errorf("code is %v: unknown", input.Code))
				return
			}
			ctx.JSON(http.StatusBadRequest, fmt.Errorf("error is %s", err.Error()))
			return
		}
		cipher, err := aes.NewAES(invitationBody.Key)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "AESの初期化に失敗"})
			return
		}
		decodedIV, err := hex.DecodeString(invitationBody.IV)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("IVデコードの失敗")})
			return
		}
		decodedET, err := hex.DecodeString(invitationBody.EncryptedText)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "Encrypted Textデコードの失敗"})
			return
		}
		// IDとパスワードが正しいか
		_, err = crypto.Decrypt(cipher, decodedIV, input.String(), decodedET)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": fmt.Sprintf("faild to decrypt. check whether password or id is wrong error is %s", err.Error())})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": invitationBody.YoutubeURL})
	}
}

// Invitation 招待コードを生成するハンドラー
func Invitation(app *fb.FirebaseApp, db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var input container.Input
		if err := ctx.BindJSON(&input); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": "faild to bind json"})
			return
		}

		//TODO 匿名ユーザの場合招待コードは生成できないようにする
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
		// dbにインサート
		if err := repository.InsertInvitationCodeWithUser(user.UID, con, db); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		ctx.JSON(http.StatusCreated, gin.H{"message": con.Code})
	}
}

// InvitationGuest 招待コードを生成するハンドラー
func InvitationGuest(db *sqlx.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
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
		if err := repository.InsertInvitationCodeWithGuest(con, db); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
			return
		}
		log.Println("DBインサート成功")
		ctx.JSON(http.StatusCreated, gin.H{"message": con.Code})
	}
}
