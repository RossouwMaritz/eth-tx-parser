package storage

type Storage interface {
	GetLastProcessedBlock() int
	SetLastProcessedBlock(block int)
	GetSubscribedAddresses() []string
	SubscribeAddress(address string)
}
