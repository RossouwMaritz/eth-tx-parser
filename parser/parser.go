package parser

import (
	"eth-tx-parser/api"
	"eth-tx-parser/storage"
	"eth-tx-parser/types"
	"fmt"
	"log/slog"
	"time"
)

type TransactionParser struct {
	storage storage.Storage
	api     api.API
}

func NewParser(storage storage.Storage, api api.API) *TransactionParser {
	return &TransactionParser{
		storage: storage,
		api:     api,
	}
}

func (tp *TransactionParser) GetCurrentBlock() int {
	return tp.storage.GetLastProcessedBlock()
}

func (tp *TransactionParser) Subscribe(address string) bool {
	tp.storage.SubscribeAddress(address)
	return true
}

func (tp *TransactionParser) GetTransactions(address string) []types.Transaction {
	currentBlock := tp.GetCurrentBlock()
	return tp.api.GetTransactions(fmt.Sprintf("0x%x", currentBlock), address)
}

func (tp *TransactionParser) Poll() {
	for {
		slog.Info("polling for transactions...")
		addresses := tp.storage.GetSubscribedAddresses()
		for _, address := range addresses {
			transactions := tp.GetTransactions(address)
			notifyTransactions(address, transactions)

		}

		currentBlock := tp.api.GetCurrentBlock()
		tp.storage.SetLastProcessedBlock(currentBlock)

		time.Sleep(time.Second * 30)
	}
}

func notifyTransactions(address string, transactions []types.Transaction) {
	for _, transaction := range transactions {
		slog.Info(
			"new transaction processed",
			slog.String("transaction_hash", transaction.Hash),
			slog.String("address", address),
		)
	}
}
