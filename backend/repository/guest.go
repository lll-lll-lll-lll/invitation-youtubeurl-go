package repository

import (
	"encoding/hex"

	"github.com/jmoiron/sqlx"
	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/container"
)

func InsertInvitationCodeWithGuestFunc(req interface{}, db *sqlx.DB) error {
	castedReq := req.(InvitationCodeWithGuest)
	_, err := db.Exec("INSERT INTO invitation_codes(code) VALUES($1)", castedReq.InvitationCode)
	if err != nil {
		return err
	}
	stmt, err := db.Prepare(`
		INSERT INTO invitation_guest (invitation_code, iv, key, encrypted_text, url)
		SELECT invitation_codes.code, $2, $3, $4, $5
		FROM invitation_codes
		WHERE invitation_codes.code = $1;`)
	if err != nil {
		return err
	}
	_, err = stmt.Exec(castedReq.InvitationCode, castedReq.IV, castedReq.Key, castedReq.EncryptedText, castedReq.YoutubeURL)
	if err != nil {
		return err
	}
	return nil
}

func InsertInvitationCodeWithGuest(con *container.Container, db *sqlx.DB) error {
	hexIV := hex.EncodeToString(con.IV)
	hexEncryptedText := hex.EncodeToString(con.EncryptedText)
	input := InvitationCodeWithGuest{InvitationCode: con.Code, IV: hexIV, Key: con.Key, EncryptedText: hexEncryptedText, YoutubeURL: con.YoutubeURL}
	if err := Transaction(db, input, InsertInvitationCodeWithGuestFunc); err != nil {
		return err
	}
	return nil
}

type InvitationCodeWithGuest struct {
	InvitationCode string `json:"invitation_code"`
	IV             string `json:"iv"`
	Key            string `json:"key"`
	EncryptedText  string `json:"encrypted_text"`
	YoutubeURL     string `json:"url"`
}
