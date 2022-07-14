package demo3

import (
	"eth_study/eth_go/demo3/ab2_t2"
	"eth_study/eth_go/demo3/nt1"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"testing"
)

func TestV1(t *testing.T) {
	// 合約地址
	// 0x2c60005875Fa4C43d0127508e6a03A6Fc97a964e
	constantA2 := common.HexToAddress("0x8f2A7a85258333de51C713c64A4F9bD97794CBbC")

	dial, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	t2, err := ab2_t2.NewAB2T2(constantA2, dial)
	if err != nil {
		log.Fatalln(err)
	}

	name, err := t2.Name(nil) //  string public name = "hallen";
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println(name)
}

func TestNT1(t *testing.T) {
	// 合約地址
	// 0xe8cc04218a3bA0da808542B39557d1c3B470b14D
	constant := common.HexToAddress("0xe8cc04218a3bA0da808542B39557d1c3B470b14D")

	dial, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	nt, err := nt1.NewNt1(constant, dial)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(nt.GetName(nil, "hello world"))
}

//err := eth_utils.InitEthClient("http://127.0.0.1:8545")
//if err != nil {
//	log.Fatalln(err)
//}
//
//list, err := eth.GetAccountList()
//if err != nil {
//	panic(err)
//}
//
//log.Println(list)

// 0xf79297b0b3e7010729f09747519981ff6c787f2a from
// 0xd3eed2b1537381e2c783298e82725fda93fda935 to

// 合約賬戶 hax => address
//mainAccount := common.HexToAddress("0xf79297b0b3e7010729f09747519981ff6c787f2a")
//toAccount := common.HexToAddress("0xd3eed2b1537381e2c783298e82725fda93fda935")

//testAbi2Obj, err := demo3_abi.NewTestAbi2(mainAccount, dial)
//if err != nil {
//	panic(err)
//}
//
//log.Println(testAbi2Obj.TestAbi2Caller)
//log.Println(testAbi2Obj.TestAbi2Transactor)
//log.Println(testAbi2Obj.TestAbi2Filterer)
//testAbi2Obj.GetMoney()
