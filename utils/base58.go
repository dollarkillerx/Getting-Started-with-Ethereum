package utils

import (
	"bytes"
	"math/big"
)

var base58 = []byte("123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz")

// Base58Encoding Base58Encoding...
func Base58Encoding(str string) string { //Base58编码
	//1. 转换成ascii码对应的值
	strByte := []byte(str)
	//2. 转换十进制
	strTen := big.NewInt(0).SetBytes(strByte)
	//3. 取出余数
	var modSlice []byte
	for strTen.Cmp(big.NewInt(0)) > 0 {
		mod := big.NewInt(0) //余数
		strTen58 := big.NewInt(58)
		strTen.DivMod(strTen, strTen58, mod)             //取余运算
		modSlice = append(modSlice, base58[mod.Int64()]) //存储余数,并将对应值放入其中
	}
	//  处理0就是1的情况 0使用字节'1'代替
	for _, elem := range strByte {
		if elem != 0 {
			break
		} else if elem == 0 {
			modSlice = append(modSlice, byte('1'))
		}
	}
	ReverseModSlice := reverseByteArr(modSlice)
	return string(ReverseModSlice)
}

// Base58Decoding 就是编码的逆过程
func Base58Decoding(str string) string { //Base58解码
	strByte := []byte(str)
	ret := big.NewInt(0)
	for _, byteElem := range strByte {
		index := bytes.IndexByte(base58, byteElem) //获取base58对应数组的下标
		ret.Mul(ret, big.NewInt(58))               //相乘回去
		ret.Add(ret, big.NewInt(int64(index)))     //相加
	}
	return string(ret.Bytes())
}

// reverseByteArr 将字节的数组反转
func reverseByteArr(bytes []byte) []byte {
	for i := 0; i < len(bytes)/2; i++ {
		bytes[i], bytes[len(bytes)-1-i] = bytes[len(bytes)-1-i], bytes[i] //前后交换
	}
	return bytes
}
