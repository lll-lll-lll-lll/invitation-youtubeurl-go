package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"fmt"
	"io"
)

func NewAES(key string) (cipher.Block, error) {
	c, err := aes.NewCipher([]byte(key))
	if err != nil {
		return nil, err
	}
	return c, nil
}

// GenerateIV 暗号化する際に使用するランダムな文字列
func GenerateIV() ([]byte, error) {
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}
	return iv, nil
}

func Encrypt(cBlock cipher.Block, iv []byte, plaintext string) []byte {
	enc := cipher.NewCFBEncrypter(cBlock, iv)
	ciphertext := make([]byte, len(plaintext))
	enc.XORKeyStream(ciphertext, []byte(plaintext))
	return ciphertext
}

func Decrypt(cBlock cipher.Block, iv []byte, plaintext string, ciphertext []byte) (plaintextCopy []byte, errRes error) {
	// 復号の際に入力文字列が正しくないとpanicが発生することがあるので、そのリカバー
	defer func() {
		if err := recover(); err != nil {
			errRes = fmt.Errorf("error: %s", err)
		}
	}()
	cfbdec := cipher.NewCFBDecrypter(cBlock, iv)
	plaintextCopy = make([]byte, len(plaintext))
	cfbdec.XORKeyStream(plaintextCopy, ciphertext)
	return plaintextCopy, nil
}

// RandomString 指定したバイト数でランダムな文字列を生成するメソッド
func RandomString(byteNum uint32) (string, error) {
	const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	b := make([]byte, byteNum)
	if _, err := rand.Read(b); err != nil {
		return "", err
	}
	var result string
	for _, v := range b {
		// index が letters の長さに収まるように調整
		result += string(letters[int(v)%len(letters)])
	}
	return result, nil
}

// GenerateKeyAndIV KeyとIVをまとめて生成するメソッド
func GenerateKeyAndIV(byteNum uint32) (string, []byte, error) {
	key, err := RandomString(uint32(byteNum))
	if err != nil {
		return "", nil, err
	}
	iv, err := GenerateIV()
	if err != nil {
		return "", nil, err
	}
	return key, iv, nil
}
