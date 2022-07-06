package utils

import (
	"log"
	"testing"
)

func TestHash(t *testing.T) {
	p := "hello world"
	log.Println(GenMD4(p))
	log.Println(GenMD5(p))
	log.Println(GenSHA256(p))
}
