package utils

import (
	"fmt"
	"log"
	"testing"
)

func TestBase58(t *testing.T) {
	src := "hello world"
	ret := Base58Encoding(src)
	fmt.Println(ret)
	decoding := Base58Decoding(ret)
	log.Println(string(decoding))
}
