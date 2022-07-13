# Geth  https://geth.ethereum.org/

### 常用命令

https://learnblockchain.cn/2017/11/29/geth_cmd_options

- `--datadir` 指定數據目錄
  - `geth --data dir "dir" account new`
- 賬戶相關
  - `geth account list` 查看賬戶
  - `geth account new 輸入兩次密碼` 創建賬戶
  - `geth account update "user addr" 輸入兩次密碼` 更新指定賬戶
  - `geth --datadir <DATADIR> account new` 指定保存路徑的創建用戶

#### RPC Port

`geth --identity "node name" --rpc --rpcport 8545 --rpcapi "db,eth,web3,personal"`

- `--identity "節點名稱"` 節點身份標識
- `--rpc` 開啓RPC接口
- `--rpcport` RPC端口,默認是go是8545
- `--rpcapi "db,eth,net,web3"` 提供給別人使用的PRC API,默認為WEB3接口
- `--rpccorsdomain` 設置能鏈接到你節點的URL 用來完成RPC任務 *值任何URL都能連接到
- `--datadir` 區塊數據文件夾(在geth同目錄下生成data0文件夾)
- `--networkid` new_version的id
- `--port` 用來監聽其他節點的端口
- `--nodiscover` 你的節點不會被其他人發現,除非他們手動添加你  

我們用下面的就行了:

`geth --identity "v1" --rpc --rpcport 8545 --rpcapi "net,personal,eth" --datadir "local_data" --http.corsdomain "*" -allow-insecure-unlock console `

目錄結構:
- local_data
  - geth        存放鏈信息
  - keystore    存放用戶信息

#### console 下的一些基礎命令

- `eth.coinbase` 
- `eth.accounts`
- `eth.getBalance("xxx")`
- `miner.start(1)` // 挖礦 綫程數
- `eth.mining`      // 是否在挖礦
- `miner.stop()`    // 停止挖礦

#### 初始化私有節點 創造世塊

``` 
{
  "config": {
    "chainId": 666,  // 鏈地址
    "homesteadBlock": 0,
    "eip150Block": 0,
    "eip155Block": 0,
    "eip158Block": 0
  },
  "nonce": "0x0",
  "timestamp": "0x00",
  "extraData": "", // 擴展數據
  "gasLimit": "0xffffffff", 
  "difficulty": "0x00002",   // 難度
  "mixHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "coinbase": "0x0000000000000000000000000000000000000000",
  "number": "0x0",
  "parentHash": "0x0000000000000000000000000000000000000000000000000000000000000000",
  "alloc": { }
}
```

初始化:

`geth -datadir "local_data2" init genesis.json`

`geth --identity "v1" --rpc --rpcport 8545 --rpcapi "net,personal,eth" --datadir "local_data2" --http.corsdomain "*" -allow-insecure-unlock console`
