package demo3

import (
	"context"
	"crypto/ecdsa"
	"eth_study/eth_go/demo3/ab2_t2"
	"eth_study/eth_go/demo3/erc20"
	"eth_study/eth_go/demo3/nt1"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"log"
	"math/big"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"golang.org/x/crypto/sha3"
)

func TestErc20(t *testing.T) {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
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

	//{
	//	// 加载您的私钥
	//	privateKey, err := crypto.HexToECDSA("c00b13047916c5dfb58944d252996c9e436ff1a9a333d53ad646a275347d80e3")
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	publicKey := privateKey.Public()
	//	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	//	if !ok {
	//		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	//	}
	//
	//	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	//	// 读取我们应该用于帐户交易的随机数
	//	nonce, err := dial.PendingNonceAt(context.Background(), fromAddress)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	// 獲取gas費用
	//	gasPrice, err := dial.SuggestGasPrice(context.Background())
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//
	//	// 獲取鏈id
	//	chainID, err := dial.NetworkID(context.Background())
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//
	//	//auth := bind.NewKeyedTransactor(privateKey)
	//	// 生成用於簽名的信息
	//	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	auth.Nonce = big.NewInt(int64(nonce))
	//	auth.Value = big.NewInt(0)     // in wei
	//	auth.GasLimit = uint64(300000) // in units
	//	auth.GasPrice = gasPrice
	//
	//	// 調用轉賬
	//	transfer, err := eRC20.Transfer(auth, common.HexToAddress("0xB1B775149F08aC25C2a797f827D6365E8123B802"), big.NewInt(100))
	//	if err != nil {
	//		log.Fatalln(err)
	//	}
	//	log.Println(transfer.ChainId().String())
	//	log.Println(transfer.Hash().String())
	//}

	{
		// 代幣專賬2

		// 加載私鑰
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

		// 代币传输不需要传输ETH，因此将交易“值”设置为“0”。
		value := big.NewInt(0)

		// 先将您要发送代币的地址存储在变量中。
		toAddress := common.HexToAddress("0xB1B775149F08aC25C2a797f827D6365E8123B802")

		//  然后我们使用函数名的keccak-256哈希来检索 方法ID，它是前8个字符（4个字节）。 然后，我们附加我们发送的地址，并附加我们打算转账的代币数量。 这些输入需要256位长（32字节）并填充左侧。 方法ID不需填充。
		tokenAddress := common.HexToAddress("0x9Ac6bd28e76Ac936D35B80ee1002265519E7e923")

		// 函数名将是传递函数的名称，即ERC-20规范中的transfer和参数类型。 第一个参数类型是address（令牌的接收者），第二个类型是uint256（要发送的代币数量）。
		transferFnSignature := []byte("transfer(address,uint256)")

		// go-ethereum导入crypto/sha3包以生成函数签名的Keccak256哈希。 然后我们只使用前4个字节来获取方法ID。
		hash := sha3.NewLegacyKeccak256()
		hash.Write(transferFnSignature)
		methodID := hash.Sum(nil)[:4]
		fmt.Println(hexutil.Encode(methodID)) // 0xa9059cbb

		// 我们需要将给我们发送代币的地址左填充到32字节。
		paddedAddress := common.LeftPadBytes(toAddress.Bytes(), 32)
		fmt.Println(hexutil.Encode(paddedAddress)) // 0x0000000000000000000000004592d8f8d7b001e72cb26a73e4fa1806a51ac79d

		// 接下来我们确定要发送多少个代币，在这个例子里是1,000个，并且我们需要在big.Int中格式化为wei。
		amount := new(big.Int)
		amount.SetString("10000000", 10) // 1000 tokens
		// 代币量也需要左填充到32个字节。
		paddedAmount := common.LeftPadBytes(amount.Bytes(), 32)
		fmt.Println(hexutil.Encode(paddedAmount)) // 0x00000000000000000000000000000000000000000000003635c9adc5dea00000
		// 接下来我们只需将方法ID，填充后的地址和填后的转账量，接到将成为我们数据字段的字节片。
		var data []byte
		data = append(data, methodID...)
		data = append(data, paddedAddress...)
		data = append(data, paddedAmount...)
		// 这个函数从ethereum包中获取CallMsg结构，我们在其中指定数据和地址。 它将返回我们估算的完成交易所需的估计燃气上限。
		gasLimit, err := dial.EstimateGas(context.Background(), ethereum.CallMsg{
			To:   &toAddress,
			Data: data,
		})
		if err != nil {
			log.Fatal(err)
		}

		gasPrice, err := dial.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		log.Println(gasLimit) // 23256
		log.Println(gasPrice)
		gasLimit = uint64(200000)
		// 构建交易事务类型，这类似于您在ETH转账部分中看到的，除了to字段将是代币智能合约地址。 这个常让人困惑。我们还必须在调用中包含0 ETH的值字段和刚刚生成的数据字节。
		tx := types.NewTransaction(nonce, tokenAddress, value, gasLimit, gasPrice, data)
		// 下一步是使用发件人的私钥对事务进行签名。 SignTx方法需要EIP155签名器(EIP155 signer)，这需要我们从客户端拿到链ID。
		chainID, err := dial.NetworkID(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		signedTx, err := types.SignTx(tx, types.NewEIP155Signer(chainID), privateKey)
		if err != nil {
			log.Fatal(err)
		}

		// 最后广播交易。
		err = dial.SendTransaction(context.Background(), signedTx)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("tx sent: %s", signedTx.Hash().Hex()) // tx sent: 0xa56316b637a94c4cc0331c73ef26389d6c097506d581073f927275e7a6ece0bc
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
