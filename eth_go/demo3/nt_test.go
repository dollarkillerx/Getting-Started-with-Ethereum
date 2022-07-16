package demo3

import (
	"context"
	"crypto/ecdsa"
	"eth_study/eth_go/demo3/ab2_t2"
	"eth_study/eth_go/demo3/erc20"
	"eth_study/eth_go/demo3/nt1"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TestErc20(t *testing.T) {
	// 參考
	// https://caohuan.tech/2021/09/06/%E4%BB%A5%E5%A4%AA%E5%9D%8A/%E4%BD%BF%E7%94%A8go%E8%B0%83%E7%94%A8%E6%99%BA%E8%83%BD%E5%90%88%E7%BA%A6/
	// https://goethereumbook.org/zh/transfer-tokens/

	// 0x9Ac6bd28e76Ac936D35B80ee1002265519E7e923
	constant := common.HexToAddress("0x9Ac6bd28e76Ac936D35B80ee1002265519E7e923")
	dial, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	eRC20, err := erc20.NewERC20(constant, dial)
	if err != nil {
		log.Fatalln(err)
	}

	{
		of, err := eRC20.BalanceOf(nil, common.HexToAddress("0x78324E5B83Ecc466B8A0F7be298712993CCb40E2"))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(of)
	}

	{
		of, err := eRC20.BalanceOf(nil, common.HexToAddress("0xB1B775149F08aC25C2a797f827D6365E8123B802"))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(of)
	}

	{
		privateKey, err := crypto.HexToECDSA("c00b13047916c5dfb58944d252996c9e436ff1a9a333d53ad646a275347d80e3")
		if err != nil {
			log.Fatal(err)
		}
		publicKey := privateKey.Public()
		publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
		if !ok {
			log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		}

		fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
		nonce, err := dial.PendingNonceAt(context.Background(), fromAddress)
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := dial.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		chainID, err := dial.NetworkID(context.Background())
		if err != nil {
			log.Fatalln(err)
		}

		//auth := bind.NewKeyedTransactor(privateKey)
		auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
		if err != nil {
			log.Fatalln(err)
		}
		auth.Nonce = big.NewInt(int64(nonce))
		auth.Value = big.NewInt(0)     // in wei
		auth.GasLimit = uint64(300000) // in units
		auth.GasPrice = gasPrice

		transfer, err := eRC20.Transfer(auth, common.HexToAddress("0xB1B775149F08aC25C2a797f827D6365E8123B802"), big.NewInt(100))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(transfer.ChainId().String())
		log.Println(transfer.Hash().String())
	}
}

func TestErc202(t *testing.T) {
	// 0xc2132d05d31c914a87c6611c10748aeb04b58e8f
	constant := common.HexToAddress("0xc2132d05d31c914a87c6611c10748aeb04b58e8f")
	dial, err := ethclient.Dial("https://polygon-rpc.com")
	if err != nil {
		panic(err)
	}

	eRC20, err := erc20.NewERC20(constant, dial)
	if err != nil {
		log.Fatalln(err)
	}

	{
		of, err := eRC20.BalanceOf(nil, common.HexToAddress("0xd1f84EC6652b315e77EbFFA649f6742754574796"))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(of)
	}

	{
		of, err := eRC20.BalanceOf(nil, common.HexToAddress("0xB2EaC6c72B225c6a24Ad08bD03D2B29a0219d9f1"))
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(of)
	}
}

func TestV1(t *testing.T) {
	// 合約地址
	// 0xAbE050045031a9A3849bf8789bd4D57610A5A05F
	constantA2 := common.HexToAddress("0xAbE050045031a9A3849bf8789bd4D57610A5A05F")

	dial, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		panic(err)
	}

	t2, err := ab2_t2.NewTest2Abb(constantA2, dial)
	if err != nil {
		log.Fatalln(err)
	}

	name, err := t2.GetName(nil, "HELLOWORLD") //  string public name = "hallen";
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
