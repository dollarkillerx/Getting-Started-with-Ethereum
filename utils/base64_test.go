package utils

import (
	"fmt"
	"log"
	"testing"
)

func TestBase64(t *testing.T) {
	src := "hello world"
	ret := Base64Encoding(src)
	fmt.Println(ret)
	decoding, err := Base64Decoding(ret)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(decoding))
}

func TestURLBase64(t *testing.T) {
	src := "hello world"
	ret := Base64URLEncoding(src)
	fmt.Println(ret)
	decoding, err := Base64URLDecoding(ret)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(decoding))
}
