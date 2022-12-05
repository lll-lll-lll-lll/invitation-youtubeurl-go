package invitation

import (
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

func New(inputCode string) (*InvitationCode, error) {
	code := Code(inputCode)
	if err := code.Validate(); err != nil {
		return nil, err
	}
	return &InvitationCode{Code: code}, nil
}
