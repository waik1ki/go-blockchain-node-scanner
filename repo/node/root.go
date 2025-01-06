package node

import (
	"go-blockchain-scope/env"
	. "go-blockchain-scope/utils"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Node struct {
	env *env.Env

	client *ethclient.Client
}

type NodeImpl interface {
	GetChainID() *big.Int
	GetClient() *ethclient.Client
	GetBlockByNumber(number *big.Int) *types.Block
	GetLatestBlock() uint64
	GetReceiptByHash(hash common.Hash) *types.Receipt
}

func NewNode(env *env.Env) (NodeImpl, error) {
	n := &Node{env: env}

	var err error

	if n.client, err = ethclient.Dial(env.Node.Dial); err != nil {
		panic(err)
	} else {
		return n, nil
	}
}

func (n *Node) GetReceiptByHash(hash common.Hash) *types.Receipt {
	if res, err := n.client.TransactionReceipt(Context(), hash); err != nil {
		log.Println(err)
		return nil
	} else {
		return res
	}
}

func (n *Node) GetLatestBlock() uint64 {
	if res, err := n.client.BlockNumber(Context()); err != nil {
		log.Println(err)
		return 0
	} else {
		return res
	}
}

func (n *Node) GetBlockByNumber(number *big.Int) *types.Block {
	if res, err := n.client.BlockByNumber(Context(), number); err != nil {
		log.Println(err)
		return nil
	} else {
		return res
	}
}

func (n *Node) GetChainID() *big.Int {
	if res, err := n.client.ChainID(Context()); err != nil {
		log.Println(err)
		return big.NewInt(0)
	} else {
		return res
	}
}

func (n *Node) GetClient() *ethclient.Client {
	return n.client
}
