package main

import (
	"eth_study/eth_go/demo3/demo3_abi"
	"eth_study/eth_go/eth_utils"
	"eth_study/eth_go/eth_utils/eth"

	"log"
)

func main() {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	list, err := eth.GetAccountList()
	if err != nil {
		panic(err)
	}

	log.Println(list)

	demo3_abi.NewTestAbi2()
}
