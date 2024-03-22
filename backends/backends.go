package backends

import (
	"eth-tx-parser/api"
	"eth-tx-parser/api/cloudflare"
	"eth-tx-parser/parser"
	"eth-tx-parser/storage"
	"eth-tx-parser/storage/memory"
)

type Backends struct {
	Storage storage.Storage
	API     api.API
	Parser  parser.Parser
}

func NewBackends() *Backends {
	storage := memory.NewStorage()
	api := cloudflare.NewAPI(storage)

	return &Backends{
		Storage: storage,
		API:     api,
		Parser:  parser.NewParser(storage, api),
	}
}
