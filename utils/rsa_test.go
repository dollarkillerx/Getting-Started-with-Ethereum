package utils

import (
	"log"
	"testing"
)

func TestGenAESKey(t *testing.T) {
	priKey, pubKey, err := GenerateRSAKey(1028)
	if err != nil {
		log.Fatalln(err)
	}

	log.Println(priKey)
	log.Println(pubKey)

	message := "hello world"

	encrypt, err := RSAEncrypt(message, pubKey)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(Base64Encoding(encrypt))

	decrypt, err := RSADecrypt(encrypt, priKey)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(decrypt))

	sign, err := RSASign(message, priKey)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(Base64Encoding(sign))

	//message += "x"
	err = RSASignVer(message, sign, pubKey)
	log.Println(err)
}
