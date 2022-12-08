package invitation

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
)

type Code string

type InvitationCode struct {
	Code Code `json:"code"`
}

func (c Code) Validate() error {
	if c == "0" || c == "" {
		return errors.New(fmt.Sprintf("code invalid. your code is %s", c))
	}
	return nil
}

// GenerateRandomCode 8文字のランダムな文字列を生成
func GenerateRandomCode() (string, error) {
	length := 6
	b := make([]byte, length)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.RawURLEncoding.EncodeToString(b), nil
}

func New(inputCode string) (*InvitationCode, error) {
	code := Code(inputCode)
	if err := code.Validate(); err != nil {
		return nil, err
	}
	return &InvitationCode{Code: code}, nil
}
