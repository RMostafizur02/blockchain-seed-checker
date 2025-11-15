package scanner

import (
	"crypto-wallet-seed-checker/pkg/config"
	"crypto-wallet-seed-checker/pkg/types"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

// ScanAddresses scans addresses across all supported blockchains
func ScanAddresses(addresses map[string][]types.WalletAddress) (map[string]config.ScanResult, error) {
	results := make(map[string]config.ScanResult)
	
	for chain, walletAddresses := range addresses {
		if len(walletAddresses) == 0 {
			continue
		}
		
		// Scan the first address for each chain (for demo purposes)
		address := walletAddresses[0]
		result, err := ScanAddress(chain, address.Address)
		if err != nil {
			results[chain] = config.ScanResult{
				Chain:   chain,
				Address: address.Address,
				Error:   err.Error(),
			}
		} else {
			results[chain] = result
		}
	}
	
	return results, nil
}

// ScanAddress scans a single address for a specific blockchain
func ScanAddress(chain, address string) (config.ScanResult, error) {
	chainConfig, exists := config.Blockchains[chain]
	if !exists {
		return config.ScanResult{}, fmt.Errorf("unsupported chain: %s", chain)
	}

	// Use the first API endpoint
	if len(chainConfig.APIs) == 0 {
		return config.ScanResult{}, fmt.Errorf("no API endpoints configured for %s", chain)
	}

	apiURL := fmt.Sprintf(chainConfig.APIs[0], address)
	
	client := &http.Client{Timeout: 30 * time.Second}
	
	resp, err := client.Get(apiURL)
	if err != nil {
		return config.ScanResult{}, fmt.Errorf("API request failed: %v", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return config.ScanResult{}, fmt.Errorf("failed to read response: %v", err)
	}

	// Parse response based on chain
	switch chain {
	case "ethereum", "bsc", "polygon":
		return parseEVMResponse(chain, address, body, chainConfig)
	case "bitcoin":
		return parseBitcoinResponse(address, body, chainConfig)
	case "dogecoin":
		return parseDogecoinResponse(address, body, chainConfig)
	default:
		return config.ScanResult{
			Chain:      chain,
			Address:    address,
			Balance:    0,
			Unit:       chainConfig.Symbol,
			HasBalance: false,
		}, nil
	}
}

func parseEVMResponse(chain, address string, body []byte, config config.BlockchainConfig) (config.ScanResult, error) {
	var response struct {
		Status  string `json:"status"`
		Message string `json:"message"`
		Result  string `json:"result"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return config.ScanResult{}, fmt.Errorf("failed to parse API response: %v", err)
	}

	if response.Status != "1" {
		return config.ScanResult{
			Chain:   chain,
			Address: address,
			Error:   response.Message,
		}, nil
	}

	balanceWei, err := strconv.ParseInt(response.Result, 10, 64)
	if err != nil {
		return config.ScanResult{}, fmt.Errorf("failed to parse balance: %v", err)
	}

	balance := float64(balanceWei) / float64(config.Decimals)
	hasBalance := balance > 0

	return config.ScanResult{
		Chain:      chain,
		Address:    address,
		Balance:    balance,
		Unit:       config.Symbol,
		HasBalance: hasBalance,
	}, nil
}

func parseBitcoinResponse(address string, body []byte, config config.BlockchainConfig) (config.ScanResult, error) {
	var response struct {
		ChainStats struct {
			FundedTxoSum int64 `json:"funded_txo_sum"`
		} `json:"chain_stats"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return config.ScanResult{}, fmt.Errorf("failed to parse Bitcoin API response: %v", err)
	}

	balanceSatoshi := response.ChainStats.FundedTxoSum
	balanceBTC := float64(balanceSatoshi) / float64(config.Decimals)
	hasBalance := balanceBTC > 0

	return config.ScanResult{
		Chain:      "bitcoin",
		Address:    address,
		Balance:    balanceBTC,
		Unit:       config.Symbol,
		HasBalance: hasBalance,
	}, nil
}

func parseDogecoinResponse(address string, body []byte, config config.BlockchainConfig) (config.ScanResult, error) {
	var response struct {
		Balance string `json:"balance"`
		Success int    `json:"success"`
	}

	if err := json.Unmarshal(body, &response); err != nil {
		return config.ScanResult{}, fmt.Errorf("failed to parse Dogecoin API response: %v", err)
	}

	if response.Success != 1 {
		return config.ScanResult{
			Chain:   "dogecoin",
			Address: address,
			Error:   "API returned error",
		}, nil
	}

	balance, err := strconv.ParseFloat(response.Balance, 64)
	if err != nil {
		return config.ScanResult{}, fmt.Errorf("failed to parse Dogecoin balance: %v", err)
	}

	return config.ScanResult{
		Chain:      "dogecoin",
		Address:    address,
		Balance:    balance,
		Unit:       config.Symbol,
		HasBalance: balance > 0,
	}, nil
}

// DeriveAllAddresses is a wrapper around the deriver package function
func DeriveAllAddresses(seed []byte) (map[string][]types.WalletAddress, error) {
	// This would typically call the deriver package
	// For now, return empty result
	return make(map[string][]types.WalletAddress), nil
}
