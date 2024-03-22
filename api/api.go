package api

import "eth-tx-parser/types"

type API interface {
	GetCurrentBlock() int
	GetTransactions(fromBlock, address string) []types.Transaction
}
