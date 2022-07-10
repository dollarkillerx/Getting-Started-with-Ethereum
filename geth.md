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

`geth --identity "v1" --rpc --rpcport 8545 --rpcapi "net,personal,eth" --datadir "local_data" -allow-insecure-unlock`

目錄結構:
- local_data
  - geth        存放鏈信息
  - keystore    存放用戶信息