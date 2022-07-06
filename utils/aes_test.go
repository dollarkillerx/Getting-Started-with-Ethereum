package utils

import (
	"log"
	"testing"
)

func TestAESEncoding(t *testing.T) {
	key := GenAESKey(AESType128)

	pc := "hello worlddsffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbghytrjyturjkuykyiutkevwvrybyukynmilmo8ilm8oln7bvtycrec4rhcrhevrtvrthv64r 57 jjbv657jb6betrvgertwcfrtfcdertcferwt"
	encoding, err := AESEncoding(key, pc)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(Base64Encoding(encoding))
	log.Println(len(Base64Encoding(encoding)))

	decrypt, err := AESDecrypt(key, encoding)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(decrypt))
}
