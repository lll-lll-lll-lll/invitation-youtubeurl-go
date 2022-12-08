package repository

import (
	"encoding/hex"

	"github.com/jmoiron/sqlx"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/container"
)

type InvitationCodeWithUser struct {
	UserID         string `json:"id"`
	InvitationCode string `json:"invitation_code"`
	IV             string `json:"iv"`
	Key            string `json:"key"`
	EncryptedText  string `json:"encrypted_text"`
	YoutubeURL     string `json:"url"`
}

func InsertInvitationCodeWithUserFunc(req interface{}, db *sqlx.DB) error {
	castedReq := req.(InvitationCodeWithUser)
	stmt, err := db.Prepare("INSERT INTO invitation(id, invitation_code, iv, key,encrypted_text, url ) VALUES($1,(SELECT invitation_codes.code FROM invitation_codes WHERE code = $2),$3,$4,$5,$6)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(castedReq.UserID, castedReq.InvitationCode, castedReq.IV, castedReq.Key, castedReq.EncryptedText, castedReq.YoutubeURL)
	if err != nil {
		return err
	}
	return nil
}

func InsertInvitationCodeWithUser(userID string, con *container.Container, db *sqlx.DB) error {
	hexIV := hex.EncodeToString(con.IV)
	hexEncryptedText := hex.EncodeToString(con.EncryptedText)
	input := InvitationCodeWithUser{UserID: userID, InvitationCode: con.Code, IV: hexIV, Key: con.Key, EncryptedText: hexEncryptedText, YoutubeURL: con.YoutubeURL}
	if err := Transaction(db, input, InsertInvitationCodeWithUserFunc); err != nil {
		return err
	}
	return nil
}

func InsertInvitationCode(con *container.Container, db *sqlx.DB) error {
	postCode := PostCode{Code: con.Code}
	if err := Transaction(db, postCode, insertCodeFunc); err != nil {
		return err
	}
	return nil
}

type PostCode struct {
	Code string `json:"code"`
}

func insertCodeFunc(req interface{}, db *sqlx.DB) error {
	castedReq := req.(PostCode)
	stmt, err := db.Prepare("INSERT INTO invitation_codes(code) VALUES($1)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(castedReq.Code)
	if err != nil {
		return err
	}
	return nil
}
