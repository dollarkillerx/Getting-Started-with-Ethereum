package utils

import (
	"crypto/md5"
	"crypto/sha256"
	"encoding/hex"
	"golang.org/x/crypto/md4"
)

// GenMD4 GenMD4
func GenMD4(src interface{}) string {
	var bts []byte
	switch r := src.(type) {
	case string:
		bts = []byte(r)
	case []byte:
		bts = r
	}

	sum := md4.New().Sum(bts)

	return hex.EncodeToString(sum)
}

// GenMD5 GenMD5
func GenMD5(src interface{}) string {
	var bts []byte
	switch r := src.(type) {
	case string:
		bts = []byte(r)
	case []byte:
		bts = r
	}

	sum := md5.New().Sum(bts)

	return hex.EncodeToString(sum)
}

// GenSHA256 GenSHA256
func GenSHA256(src interface{}) string {
	var bts []byte
	switch r := src.(type) {
	case string:
		bts = []byte(r)
	case []byte:
		bts = r
	}

	sum := sha256.New().Sum(bts)
	return hex.EncodeToString(sum)
}
