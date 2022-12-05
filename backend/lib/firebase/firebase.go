package firebase

import (
	"context"
	"fmt"
	"os"

	firebase "firebase.google.com/go"
	"firebase.google.com/go/auth"
	"google.golang.org/api/option"
)

// InitFireBase Firebaseの初期化
func InitFireBase() (*FirebaseApp, error) {
	opt := option.WithCredentialsJSON([]byte(os.Getenv("GOOGLE_CREDENTIALS_JSON")))
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
