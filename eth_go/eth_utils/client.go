package eth_utils

import (
	"github.com/ethereum/go-ethereum/rpc"
)

var Client *rpc.Client

func InitEthClient(addr string) error {
	dial, err := rpc.Dial(addr)
	//dial, err := rpc.Dial("http://127.0.0.1:8545")
	if err != nil {
		return err
	}

	Client = dial

	return nil
}
