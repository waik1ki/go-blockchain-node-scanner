package node

import (
	"go-blockchain-scope/env"

	"github.com/ethereum/go-ethereum/ethclient"
)

type Node struct {
	env *env.Env

	client *ethclient.Client
}

type NodeImpl interface{}

func NewNode(env *env.Env) (NodeImpl, error) {
	n := &Node{env: env}

	var err error

	if n.client, err = ethclient.Dial(env.Node.Dial); err != nil {
		panic(err)
	} else {
		return n, nil
	}
}
