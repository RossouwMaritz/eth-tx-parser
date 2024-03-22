package memory

type Storage struct {
	LastProcessedBlock  int
	LastFetchedAddress  string
	SubscribedAddresses []string
}

func NewStorage() *Storage {
	return &Storage{}
}

func (s *Storage) GetLastProcessedBlock() int {
	return s.LastProcessedBlock
}

func (s *Storage) SetLastProcessedBlock(lastProcessedBlock int) {
	s.LastProcessedBlock = lastProcessedBlock
}

func (s *Storage) GetSubscribedAddresses() []string {
	return s.SubscribedAddresses
}

func (s *Storage) SubscribeAddress(address string) {
	s.SubscribedAddresses = append(s.SubscribedAddresses, address)
}
