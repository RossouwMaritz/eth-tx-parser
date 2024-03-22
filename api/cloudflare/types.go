package cloudflare

type BlockNumberRequest struct {
	JSONRPC string `json:"jsonrpc"`
	Method  string `json:"method"`
	ID      string `json:"id,omitempty"`
}

type BlockNumberResult struct {
	JSONRPC string `json:"jsonrpc"`
	Result  string `json:"result"`
	ID      string `json:"id"`
}

type TraceFilterRequest struct {
	JSONRPC string            `json:"jsonrpc"`
	Method  string            `json:"method"`
	Params  TraceFilterParams `json:"method"`
	ID      string            `json:"id,omitempty"`
}

type TraceFilterParams struct {
	Count       int      `json:"count"`
	FromBlock   string   `json:"fromBlock"`
	FromAddress []string `json:"FromAddress"`
}

type TraceFilterResult struct {
	JSONRPC string   `json:"jsonrpc"`
	Result  []Result `json:"result"`
}

type Result struct {
	TransactionHash string `json:"transactionHash"`
}
