package block

import (
	"github.com/hy-u-3a4em/rana/core/source/account"
	"github.com/hy-u-3a4em/rana/core/source/block_difficulty"
	"github.com/hy-u-3a4em/rana/core/source/block_height"
	"github.com/hy-u-3a4em/rana/core/source/hash"
	"github.com/hy-u-3a4em/rana/core/source/transaction"
)

type Block struct {
	height         block_height.BlockHeight
	prevBlockHash  hash.Hash
	transactions   []transaction.Transaction
	lessor         account.Account
	difficulty     block_difficulty.BlockDifficulty
	generationHash hash.Hash
}
