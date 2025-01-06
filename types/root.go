package types

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type CustomBlock struct {
	Hash       common.Hash    `json:"hash"`
	ParentHash common.Hash    `json:"parentHash"`
	Miner      common.Address `json:"miner"`
	Root       common.Hash    `json:"root"`

	Time        uint64 `json:"time"`
	BlockNumber uint64 `json:"blockNumber"`
	Size        uint64 `json:"size"`

	ChainID int64 `json:"chainID"`
}

func MakeCustomBlock(block *types.Block, chainID int64) *CustomBlock {
	return &CustomBlock{
		Hash:       block.Hash(),
		ParentHash: block.ParentHash(),
		Miner:      block.Coinbase(),
		Root:       block.Root(),

		Time:        block.Time(),
		BlockNumber: block.NumberU64(),
		Size:        block.Size(),

		ChainID: chainID,
	}
}

type CustomTx struct {
	Tx   common.Hash     `json:"tx"`
	From common.Address  `json:"from"`
	To   *common.Address `json:"to"`

	BlockNumber *big.Int `json:"blockNumber"`
	Fee         *big.Int `json:"fee"`

	Amount string `json:"amount"`
	Size   uint64 `json:"size"`
	Nonce  uint64 `json:"nonce"`
	Time   int64  `json:"time"`
}

func MakeCustomTx(
	transaction *types.Transaction,
	receipt *types.Receipt,
	header *types.Header,
	signer types.EIP155Signer,
) *CustomTx {
	tx := &CustomTx{
		Tx:          receipt.TxHash,
		To:          transaction.To(),
		BlockNumber: header.Number,
		Fee:         new(big.Int).Mul(transaction.GasPrice(), big.NewInt(int64(receipt.GasUsed))),
		Amount:      transaction.Value().String(),
		Size:        transaction.Size(),
		Nonce:       transaction.Nonce(),
		Time:        transaction.Time().Unix(),
	}

	tx.From, _ = types.Sender(signer, transaction)

	return tx
}
