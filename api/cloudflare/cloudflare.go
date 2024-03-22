package cloudflare

import (
	"bytes"
	"encoding/json"
	"eth-tx-parser/storage"
	"eth-tx-parser/types"
	"fmt"
	"io"
	"log/slog"
	"net/http"
	"strconv"
)

var url = "https://cloudflare-eth.com"

type API struct {
	HTTPClient *http.Client
	Storage    storage.Storage
}

// NewAPI creates a new implementation of the cloudflare API
func NewAPI(storage storage.Storage) *API {
	return &API{
		HTTPClient: &http.Client{},
		Storage:    storage,
	}
}

func (a *API) GetCurrentBlock() int {
	req := BlockNumberRequest{
		JSONRPC: "2.0",
		Method:  "eth_blockNumber",
		ID:      "1",
	}

	payload, err := json.Marshal(req)
	if err != nil {
		slog.Error("failed to marshal get current block request")
		return 0
	}

	response, err := makeRequest(a.HTTPClient, "POST", bytes.NewBuffer(payload))
	if err != nil {
		slog.Error(
			"failed to make get current block http request",
			slog.Any("error", err),
		)
		return 0
	}

	result := BlockNumberResult{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		slog.Error(
			"failed to unmarshal get current block result",
			slog.Any("error", err),
		)
		return 0
	}

	currentBlock, err := strconv.ParseInt(result.Result[2:], 16, 0)
	if err != nil {
		slog.Error(
			"failed to convert response from hex into int",
			slog.Any("error", err),
		)
		return 0
	}

	return int(currentBlock)
}

func (a *API) GetTransactions(fromBlock, address string) []types.Transaction {
	req := TraceFilterRequest{
		JSONRPC: "2.0",
		Method:  "trace_filter",
		Params: TraceFilterParams{
			Count:       200,
			FromBlock:   fromBlock,
			FromAddress: []string{address},
		},
	}

	payload, err := json.Marshal(req)
	if err != nil {
		slog.Error("failed to marshal get transactions request")
		return nil
	}

	response, err := makeRequest(a.HTTPClient, "POST", bytes.NewBuffer(payload))
	if err != nil {
		slog.Error(
			"failed to make get transactions http request",
			slog.Any("error", err),
		)
		return nil
	}

	result := TraceFilterResult{}
	err = json.Unmarshal(response, &result)
	if err != nil {
		slog.Error(
			"failed to unmarshal trace filter result",
			slog.Any("error", err),
		)
		return nil
	}

	transactions := []types.Transaction{}
	for _, r := range result.Result {
		transactions = append(transactions, types.Transaction{
			Hash: r.TransactionHash,
		})
	}

	return transactions
}

// makeRequest is a helper function that makes HTTP requests to the Cloudflare API
func makeRequest(client *http.Client, method string, payload io.Reader) ([]byte, error) {
	req, err := http.NewRequest(method, url, payload)
	if err != nil {
		return nil, err
	}

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != 200 {
		return nil, fmt.Errorf("http request failed. status code: %d", res.StatusCode)
	}

	return body, nil
}
