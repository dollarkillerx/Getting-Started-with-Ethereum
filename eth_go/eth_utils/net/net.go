package net

import "eth_study/eth_go/eth_utils"

// GetNetworkID 獲取網絡ID
func GetNetworkID() (networkID string, err error) {
	err = eth_utils.Client.Call(&networkID, "net_version")

	// network id
	// "1": Mainnet
	// "2": Morden Testnet
	// "3": Ropsten Testnet
	// "4": Rinkeby Testnet
	// "42": Kovan Testnet
	return
}

// GetNetIsListening 客戶端是否處於監聽狀態
func GetNetIsListening() (isListening bool, err error) {
	err = eth_utils.Client.Call(&isListening, "net_listening")
	return
}

// GetNetPeerCount 所連接對端節點的數量
func GetNetPeerCount() (count string, err error) {
	err = eth_utils.Client.Call(&count, "net_peerCount")
	return
}
