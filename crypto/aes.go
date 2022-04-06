package crypto

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"

	"errors"
	"io"
)

const defaultAESSecret = `0M}u-vfN4VO]\v=d1@DI(oBfcw/<Ap3b9{:nP~Kp&u[se85TD#&-0m"@j#U"^7_`

// AesEncrypt input
func AesEncrypt(input []byte, secret []byte) ([]byte, error) {
	if len(input) == 0 {
		return nil, ErrInvalidData
	}
	return aesEncrypt(input, secret)
}

// AesDecrypt input
func AesDecrypt(input []byte, secret []byte) ([]byte, error) {
	if len(input) == 0 {
		return nil, ErrInvalidData
	}
	return aesDecrypt(input, secret)
}

// aesEncrypt Aes 加密
func aesEncrypt(plaintext []byte, secret []byte) ([]byte, error) {
	key := secret
	if len(key) == 0 {
		key = []byte(defaultAESSecret)
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	return gcm.Seal(nonce, nonce, plaintext, nil), nil
}

// aesDecrypt Aes 解密
func aesDecrypt(ciphertext []byte, secret []byte) ([]byte, error) {
	key := secret
	if len(key) == 0 {
		key = []byte(defaultAESSecret)
	}
	c, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {
		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(ciphertext) < nonceSize {
		return nil, errors.New("ciphertext too short")
	}

	nonce, ciphertext := ciphertext[:nonceSize], ciphertext[nonceSize:]
	return gcm.Open(nil, nonce, ciphertext, nil)
}
