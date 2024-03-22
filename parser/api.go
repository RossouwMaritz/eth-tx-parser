package parser

import "eth-tx-parser/types"

type Parser interface {
	// last parsed block
	GetCurrentBlock() int

	// add address to observer
	Subscribe(address string) bool

	// list of inbound or outbound transactions for an address
	GetTransactions(address string) []types.Transaction

	// Poll for latest transactions
	Poll()
}
