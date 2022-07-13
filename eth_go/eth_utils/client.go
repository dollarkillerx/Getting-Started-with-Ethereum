package eth_utils

import (
	"github.com/ethereum/go-ethereum/rpc"
)

var Client *rpc.Client

func InitEthClient(addr string) error {
	dial, err := rpc.Dial(addr)
	if err != nil {
		return err
	}

	Client = dial

	return nil
}
