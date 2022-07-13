package test

import (
	"eth_study/eth_go/eth_utils"
	"eth_study/eth_go/eth_utils/enum"
	"eth_study/eth_go/eth_utils/eth"
	"eth_study/eth_go/eth_utils/net"
	"eth_study/eth_go/eth_utils/personal"

	"log"
	"testing"
)

func TestEthUtilsNetworkID(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := net.GetNetworkID()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsGetNetIsListening(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := net.GetNetIsListening()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsGetNetPeerCount(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := net.GetNetPeerCount()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsGetAccountList(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := eth.GetAccountList()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsGetAccountList2(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := personal.GetAccountList()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsNewAccount(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := personal.NewAccount("1008611")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)

	// 0xf79297b0b3e7010729f09747519981ff6c787f2a 1234567
	// 0xd3eed2b1537381e2c783298e82725fda93fda935 1234567
	// 0xb89a820913a4cae5a1cc6d292a80765eedeadda1 1008611
}

func TestEthUtilsGetAccountBalance(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := eth.GetAccountBalance("0xf004e78a16e8ff68433a3f5a86f0b34f25a5c539", enum.Latest)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsGetGasPrice(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := eth.GetGasPrice()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsGetEthCoinbase(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := eth.GetCoinbase()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsGetMining(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := eth.GetMining()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsGetHashRate(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := eth.GetHashRate()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsGetTransactionCount(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := eth.GetTransactionCount("0x9a69cf14f89dcab2cb812209a69a2095aecad1d2", enum.Latest)
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsGetBlockNumber(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := eth.GetBlockNumber()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsLockAccount(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := personal.LockAccount("0xf004e78a16e8ff68433a3f5a86f0b34f25a5c539")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

func TestEthUtilsUnlockAccount(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	id, err := personal.UnlockAccount("0xf79297b0b3e7010729f09747519981ff6c787f2a", "1234567")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(id)
}

/*func TestPushString(t *testing.T) {
	err := eth_utils.InitEthClient("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalln(err)
	}

	ok, err := db.PutString("putStr", "key1", "val1")
	if err != nil {
		log.Fatalln(err)
	}
	log.Println(ok)
}
*/
