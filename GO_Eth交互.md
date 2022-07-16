# Go Eth 智能合約交互

### ABI 生成GO文件

``` 
abigen --abi xx.abi --pkg packageName --type structName --out xx.go
```

- `abi` 文件在 `remix` 部署時可以得到
- `pkg` 指定輸出文件的包名，也就是 package 名稱
- `type` 指定合約結構體名稱
- `out` 指定輸出go文件名稱

### ERC20 轉賬方式1:

通過 erc20.abi 實現

``` 
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
	
		auth := bind.NewKeyedTransactor(privateKey)
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
```