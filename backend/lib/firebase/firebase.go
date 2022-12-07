package firebase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator"
	"google.golang.org/api/option"
)

var validate = validator.New()

// InitFireBase Firebaseの初期化
func InitFireBase() (*FirebaseApp, error) {
	opt := option.WithCredentialsFile("./account.json")
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		return nil, fmt.Errorf("error initializing app: %v", err)
	}
	client, err := app.Auth(context.Background())
	if err != nil {
		return nil, fmt.Errorf("error getting Auth client: %v\n", err)
	}
	return &FirebaseApp{app, client}, nil
}

type FirebaseApp struct {
	App    *firebase.App
	Client *auth.Client
}

func (fire *FirebaseApp) VerifyIDToken(ctx context.Context, idToken string) (*auth.Token, error) {

	token, err := fire.Client.VerifyIDToken(ctx, idToken)
	if err != nil {
		return nil, fmt.Errorf("error verifying ID token: %v\n", err)
	}

	return token, nil
}

func (fire *FirebaseApp) GetUser(ctx context.Context, uid string) (*auth.UserRecord, error) {
	user, err := fire.Client.GetUser(ctx, uid)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func CreateUserWithUID(ctx context.Context, client *auth.Client, req *RegisterUser) (*auth.UserRecord, error) {
	params := (&auth.UserToCreate{}).
		UID(req.ID).
		Email(req.Email).
		EmailVerified(true).
		Password(req.Password)
	u, err := client.CreateUser(ctx, params)
	if err != nil {
		return nil, err
	}
	log.Printf("Successfully created user: %v\n", u)
	return u, nil
}

type RegisterUserBody struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

type RegisterUser struct {
	ID       string `json:"id" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Name     string `json:"name" validate:"required"`
}

// FireBaseAuthRequired ユーザの認証用ミドルウェア
func FireBaseAuthRequired(firebaseApp *FirebaseApp) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if !strings.HasPrefix(ctx.Request.Header.Get("Authorization"), "Bearer ") {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("not Set Authorization Token. set `Authorization: Bearer {Your Token}` on header")})
			ctx.Abort()
			return
		}
		idToken := strings.Split(ctx.Request.Header.Get("Authorization"), " ")[1]
		if idToken == "" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": errors.New("don't have auth token")})
			ctx.Abort()
			return
		}
		token, err := firebaseApp.VerifyIDToken(context.Background(), idToken)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error-verify-token": err.Error()})
			ctx.Abort()
			return
		}
		ctx = SetTokenContext(ctx, token)
		// 　ルーティング前に処理したいこと
		ctx.Next()
		//  ルーティング後に処理したいこと
	}
}

// SetTokenContext ユーザのIDトークンをコンテキストに保存
func SetTokenContext(ctx *gin.Context, token *auth.Token) *gin.Context {
	ctx.Set("authToken", token)
	return ctx
}

// GetTokenContext コンテキストからセットしたトークンを取得
func GetTokenContext(ctx *gin.Context) (*auth.Token, error) {
	fmt.Print(ctx.Keys)
	v, exist := ctx.Get("authToken")
	if !exist {
		return nil, errors.New("not auth token")
	}
	token := v.(*auth.Token)
	return token, nil
}
