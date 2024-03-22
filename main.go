package main

import (
	"eth-tx-parser/backends"
)

func main() {
	b := backends.NewBackends()

	// Subscribe to example address
	// TODO: Add endpoint to better manage addresses
	b.Storage.SubscribeAddress("0xEdC763b3e418cD14767b3Be02b667619a6374076")

	b.Parser.Poll()
}
