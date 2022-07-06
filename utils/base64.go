package utils

import "encoding/base64"

// Base64Encoding 編碼
func Base64Encoding(str interface{}) string {
	var bts []byte
	switch r := str.(type) {
	case string:
		bts = []byte(r)
	case []byte:
		bts = r
	}

	return base64.StdEncoding.EncodeToString(bts)
}

// Base64Decoding 解碼
func Base64Decoding(src string) ([]byte, error) {
	return base64.StdEncoding.DecodeString(src)
}

// Base64URLEncoding 編碼
func Base64URLEncoding(src interface{}) string {
	var bts []byte
	switch r := src.(type) {
	case string:
		bts = []byte(r)
	case []byte:
		bts = r
	}

	return base64.URLEncoding.EncodeToString(bts)
}

// Base64URLDecoding 解碼
func Base64URLDecoding(src string) ([]byte, error) {
	return base64.URLEncoding.DecodeString(src)
}
