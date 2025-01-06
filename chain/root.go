package chain

import (
	"go-blockchain-scope/env"
	"go-blockchain-scope/repo"
	"log"
	"math/big"
	"sync/atomic"
	"time"

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

	startBlock := start

	if end != 0 {
		// start ~ end
		c.readBlock(start, end)
	} else {
		// start ~ 최신 블록
		for {
			time.Sleep(3 * time.Second)

			latestBlock := c.getLatestBlock()

			if latestBlock == uint64(0) {
				log.Println("Failed To Get LatestBlock")
			} else if latestBlock < startBlock {
				log.Println("StartBlock Over LatestBlock")
			} else {
				go c.readBlock(startBlock, latestBlock)
				atomic.StoreUint64(&startBlock, latestBlock)
			}
		}
	}
}

func (c *Chain) readBlock(start, end uint64) {
	for i := start; i <= end; i++ {
		if blockToRead := c.getBlockByNumber(big.NewInt(int64(i))); blockToRead == nil {
			log.Println("Failed To Get Block", i)
			continue
		} else if blockToRead.Transactions().Len() == 0 {
			log.Println("Debug Transactions Len Zero", i)
			continue
		} else {
			log.Println("Scan Block Success Save Block & Tx", i)

			go c.saveBlock(blockToRead)
			go c.saveTx(blockToRead, blockToRead.Transactions().Len(), blockToRead.Header())
		}
	}
}

func (c *Chain) saveBlock(block *types.Block) {

}

func (c *Chain) saveTx(block *types.Block, len int, header *types.Header) {

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
