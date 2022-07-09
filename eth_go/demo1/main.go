package main

import (
	"github.com/ethereum/go-ethereum/rpc"

	"log"
)

// 鏈接服務
func main() {
	dial, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	// call
	var ret string

	err = dial.Call(&ret, "net_version")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ret)
}
