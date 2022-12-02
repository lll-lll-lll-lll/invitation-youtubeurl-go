package container

import (
	"crypto/cipher"
	"fmt"

	"github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/config"
	aes "github.com/lll-lll-lll-lll/youtube-url-converter-backend/lib/crypto"
)

type Input struct {
	ID       string `json:"id"`
	Password string `json:"password"`
	URL      string `json:"youtube_url"`
}

func (i Input) String() string {
	return fmt.Sprintf("%s%s", i.ID, i.Password)
}

type Container struct {
	YoutubeURL    config.YoutubeURL `json:"youtube_url"`
	IV            string            `json:"iv"`
	Key           string            `json:"key"`
	Cipher        cipher.Block
	EncryptedText []byte
}

// New youtubeのurlと、idとパスワードから生成した暗号文を持つコンテナを生成
func New(input Input) (*Container, error) {
	byteNum := 32
	plaintext := input.String()
	rawurl := input.URL
	youTubeURL, err := config.ToYouTubeURL(rawurl)
	if err != nil {
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
		IV:            string(iv),
		Key:           key,
		Cipher:        cipher,
		EncryptedText: encryptedText,
		YoutubeURL:    youTubeURL,
	}
	return container, nil
}
