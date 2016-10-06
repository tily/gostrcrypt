package gostrcrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

type StrCrypt struct {
	Key string
}

func (sc *StrCrypt) Encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher([]byte(sc.Key))
	if err != nil {
		return "", err
	}

	encrypted := make([]byte, aes.BlockSize+len(plaintext))

	iv := encrypted[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(encrypted[aes.BlockSize:], []byte(plaintext))

	return hex.EncodeToString(encrypted), nil
}

func (sc *StrCrypt) Decrypt(encrypted string) (string, error) {
	block, err := aes.NewCipher([]byte(sc.Key))
	if err != nil {
		return "", err
	}

	encryptedBytes, err := hex.DecodeString(encrypted)
	if err != nil {
		return "", err
	}

	iv := encryptedBytes[:aes.BlockSize]

	encryptedBytes = encryptedBytes[aes.BlockSize:]
	stream := cipher.NewCFBDecrypter(block, iv)
	decryptedBytes := make([]byte, len(encryptedBytes))
	stream.XORKeyStream(decryptedBytes, encryptedBytes)

	return string(decryptedBytes), nil
}
