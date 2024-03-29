package container

import (
	"fmt"

	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/config"
	inv "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/container/invitation"
	aes "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/crypto"
)

type IVType []byte

func (i *IVType) ToHexString() string {
	return fmt.Sprintf("%x", i)
}

type EncryptedTextType []byte

func (e *EncryptedTextType) ToHexString() string {
	return fmt.Sprintf("%x", e)
}

type Input struct {
	ID       string `json:"id" binding:"required"`
	Password string `json:"password" binding:"required"`
	URL      string `json:"youtube_url" binding:"required"`
}

func (i Input) String() string {
	return fmt.Sprintf("%s.%s", i.ID, i.Password)
}

type Container struct {
	// 復号に使うIV
	IV IVType `json:"iv"`
	// 復号に使うkey
	Key string `json:"key"`
	// idとパスワードとYoutubeURLを含んだ暗号文
	EncryptedText EncryptedTextType `json:"encrypted_text"`
	// 招待コード
	Code string `json:"code"`
	// youtube url
	YoutubeURL string `json:"youtube_url"`
}

// New idとパスワードから生成した暗号文を持つコンテナを生成
func New(input Input) (*Container, error) {
	byteNum := 32
	plaintext := input.String()
	rawurl := input.URL
	code, err := inv.GenerateRandomCode()
	// youtubeのURLかどうかチェック
	if err := config.ToYouTubeURL(rawurl).Validate(); err != nil {
		return nil, err
	}
	//暗号化の際に使用するkeyとivを生成
	key, iv, err := aes.GenerateKeyAndIV(uint32(byteNum))
	if err != nil {
		return nil, err
	}
	cipher, err := aes.NewAES(key)
	if err != nil {
		return nil, err
	}
	encryptedText := aes.Encrypt(cipher, iv, plaintext)
	container := &Container{
		IV:            iv,
		Key:           key,
		EncryptedText: encryptedText,
		Code:          code,
		YoutubeURL:    input.URL,
	}
	return container, nil
}
