package utils

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
)

/**
# 對稱加密

- AES-128 KEY LEN 16
- AES-192 KEY LEN 24
- AES-256 KEY LEN 32
*/

type AESType string

const (
	AESType128 AESType = "aes-128"
	AESType192 AESType = "aes-192"
	AESType256 AESType = "aes-256"
)

// GenAESKey GenAESKey
func GenAESKey(aesType AESType) string {
	switch aesType {
	case AESType128:
		return RandKey(16)
	case AESType192:
		return RandKey(24)
	case AESType256:
		return RandKey(32)
	}

	return RandKey(16)
}

// AESEncoding Encrypt string to base64 crypto using AES
func AESEncoding(key interface{}, message interface{}) ([]byte, error) {
	var keyBytes []byte
	switch r := key.(type) {
	case string:
		keyBytes = []byte(r)
	case []byte:
		keyBytes = r
	}

	var messageBytes []byte
	switch r := message.(type) {
	case string:
		messageBytes = []byte(r)
	case []byte:
		messageBytes = r
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, block.BlockSize()+len(messageBytes))
	iv := ciphertext[:block.BlockSize()]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[block.BlockSize():], messageBytes)

	return ciphertext, nil
}

// AESDecrypt Decrypt from base64 to decrypted string
func AESDecrypt(key interface{}, ciphertext []byte) ([]byte, error) {
	var keyBytes []byte
	switch r := key.(type) {
	case string:
		keyBytes = []byte(r)
	case []byte:
		keyBytes = r
	}
	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < block.BlockSize() {
		return nil, errors.New("ciphertext < block.BlockSize()")
	}
	iv := ciphertext[:block.BlockSize()]
	ciphertext = ciphertext[block.BlockSize():]

	stream := cipher.NewCFBDecrypter(block, iv)

	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
