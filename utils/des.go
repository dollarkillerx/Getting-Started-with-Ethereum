package utils

import (
	"crypto/cipher"
	"crypto/des"
	"crypto/rand"
	"errors"
	"io"
)

/*
# Des

*/

// DESEncoding DESEncoding...
func DESEncoding(key interface{}, message interface{}) ([]byte, error) {
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

	if len(keyBytes) != 8 {
		return nil, des.KeySizeError(len(keyBytes))
	}

	block, err := des.NewCipher(keyBytes)
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

// DESDecrypt DESDecrypt...
func DESDecrypt(key interface{}, ciphertext []byte) ([]byte, error) {
	var keyBytes []byte
	switch r := key.(type) {
	case string:
		keyBytes = []byte(r)
	case []byte:
		keyBytes = r
	}

	if len(keyBytes) != 8 {
		return nil, des.KeySizeError(len(keyBytes))
	}

	block, err := des.NewCipher(keyBytes)
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

// DES3DEncoding DES3DEncoding...
func DES3DEncoding(key interface{}, message interface{}) ([]byte, error) {
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

	if len(keyBytes) != 24 {
		return nil, des.KeySizeError(len(keyBytes))
	}

	block, err := des.NewTripleDESCipher(keyBytes)
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

// DES3DDecrypt DES3DDecrypt...
func DES3DDecrypt(key interface{}, ciphertext []byte) ([]byte, error) {
	var keyBytes []byte
	switch r := key.(type) {
	case string:
		keyBytes = []byte(r)
	case []byte:
		keyBytes = r
	}

	if len(keyBytes) != 24 {
		return nil, des.KeySizeError(len(keyBytes))
	}

	block, err := des.NewTripleDESCipher(keyBytes)
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
