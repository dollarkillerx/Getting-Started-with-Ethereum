package personal

import "eth_study/eth_go/eth_utils"

// GetAccountList 獲取賬戶列表
func GetAccountList() (accounts []string, err error) {
	err = eth_utils.Client.Call(&accounts, "personal_listAccounts")
	return
}

// NewAccount 創建用戶
func NewAccount(pwd string) (account string, err error) {
	err = eth_utils.Client.Call(&account, "personal_newAccount", pwd)
	return
}

// LockAccount 鎖定賬戶
func LockAccount(account string) (isLock bool, err error) {
	err = eth_utils.Client.Call(&isLock, "personal_lockAccount", account)
	return
}

// UnlockAccount 解除用戶鎖定
func UnlockAccount(account string, password string) (isUnlock bool, err error) {
	err = eth_utils.Client.Call(&isUnlock, "personal_unlockAccount", account, password)
	return
}
