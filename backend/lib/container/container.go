package container

import (
	"fmt"

	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/config"
	inv "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/container/invitation"
	aes "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/crypto"
)

type Input struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	URL      string `json:"youtube_url"`
}

func (i Input) String() string {
	return fmt.Sprintf("%s.%s.%s", i.ID, i.Password, i.URL)
}

type Container struct {
	// 復号に使うIV
	IV string `json:"iv"`
	// 復号に使うkey
	Key string `json:"key"`
	// idとパスワードとYoutubeURLを含んだ暗号文
	EncryptedText string `json:"encrypted_text"`
	// 招待コード
	Code string `json:"code"`
}

// New youtubeのurlと、idとパスワード、youtubeurlから生成した暗号文を持つコンテナを生成
func New(input Input) (*Container, error) {
	byteNum := 32
	plaintext := input.String()
	rawurl := input.URL
	code, err := inv.GenerateRandomCode()
	if err := config.ToYouTubeURL(rawurl).Validate(); err != nil {
		return nil, err
	}
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
		IV:            fmt.Sprintf("%x", iv),
		Key:           key,
		EncryptedText: fmt.Sprintf("%x", encryptedText),
		Code:          code,
	}
	return container, nil
}
