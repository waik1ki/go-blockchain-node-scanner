package chain

import (
	"go-blockchain-scope/env"
	"go-blockchain-scope/repo"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Chain struct {
	env *env.Env

	chainID *big.Int
	signer  types.EIP155Signer

	repo *repo.Repo
}

func ScanBlock(env *env.Env, repo *repo.Repo, startBlock, endBlock uint64) {
	c := &Chain{env: env, repo: repo}

	var err error

	if c.chainID = c.getChainID(); err != nil {
		panic(err)
	} else {
		c.signer = types.NewEIP155Signer(c.chainID)
		c.scanBlock(startBlock, endBlock)
	}
}

func (c *Chain) scanBlock(start, end uint64) {
	for {
	}
}

func (c *Chain) getChainID() *big.Int {
	return c.repo.Node.GetChainID()
}

func (c *Chain) getLatestBlock() uint64 {
	return c.repo.Node.GetLatestBlock()
}

func (c *Chain) getBlockByNumber(number *big.Int) *types.Block {
	return c.repo.Node.GetBlockByNumber(number)
}

func (c *Chain) getClient() *ethclient.Client {
	return c.repo.Node.GetClient()
}
