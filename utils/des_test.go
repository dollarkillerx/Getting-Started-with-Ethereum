package utils

import (
	"log"
	"testing"
)

func TestDESEncoding(t *testing.T) {
	pc := "hello worlddsffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffbghytrjyturjkuykyiutkevwvrybyukynmilmo8ilm8oln7bvtycrec4rhcrhevrtvrthv64r 57 jjbv657jb6betrvgertwcfrtfcdertcferwt"
	key := RandKey(8)
	encoding, err := DESEncoding(key, pc)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(Base64Encoding(encoding))
	log.Println(len(Base64Encoding(encoding)))

	decrypt, err := DESDecrypt(key, encoding)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(string(decrypt))
}
