package eth

import (
	"eth_study/eth_go/eth_utils"
	"eth_study/eth_go/eth_utils/enum"
)

// GetAccountList 獲取賬戶列表
func GetAccountList() (accounts []string, err error) {
	err = eth_utils.Client.Call(&accounts, "eth_accounts")
	return
}

// GetAccountBalance 獲取賬戶餘額
func GetAccountBalance(account string, param enum.BlockParameter) (balance string, err error) {
	err = eth_utils.Client.Call(&balance, "eth_getBalance", account, param)
	return
}

// GetGasPrice 獲取gas價格
func GetGasPrice() (gasPrice string, err error) {
	err = eth_utils.Client.Call(&gasPrice, "eth_gasPrice")
	return
}

// GetCoinbase 獲取挖礦賬戶地址
func GetCoinbase() (coinbase string, err error) {
	err = eth_utils.Client.Call(&coinbase, "eth_coinbase")
	return
}

// GetMining 獲取是否正在挖礦
func GetMining() (isMining bool, err error) {
	err = eth_utils.Client.Call(&isMining, "eth_mining")
	return
}

// GetHashRate 返回節點挖礦時每秒可以算出的哈希數
func GetHashRate() (hashRate string, err error) {
	err = eth_utils.Client.Call(&hashRate, "eth_hashrate")
	return
}

// GetTransactionCount 返回指定地址發生的交易數量
func GetTransactionCount(account string, param enum.BlockParameter) (count string, err error) {
	err = eth_utils.Client.Call(&count, "eth_getTransactionCount", account, param)
	return
}

// GetBlockNumber 獲取當前節點Number
func GetBlockNumber() (blockNumber string, err error) {
	err = eth_utils.Client.Call(&blockNumber, "eth_blockNumber")
	return
}

//// GetProtocolVersion 獲取eth協議版本  (廢棄)
//func GetProtocolVersion() (version string, err error) {
//	err = eth_utils.Client.Call(&version, "eth_protocolVersion")
//	return
//}
