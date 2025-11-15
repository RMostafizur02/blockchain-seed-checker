package config

type BlockchainConfig struct {
	Name            string
	DerivationPaths []string
	APIs            []string
	Symbol          string
	Decimals        int
}

var Blockchains = map[string]BlockchainConfig{
	"bitcoin": {
		Name: "Bitcoin",
		DerivationPaths: []string{
			"m/44'/0'/0'/0/0", // Legacy
			"m/84'/0'/0'/0/0", // Native SegWit
		},
		APIs: []string{
			"https://blockstream.info/api/address/%s",
			"https://blockchain.info/rawaddr/%s",
		},
		Symbol:   "BTC",
		Decimals: 8,
	},
	"ethereum": {
		Name:            "Ethereum",
		DerivationPaths: []string{"m/44'/60'/0'/0/0"},
		APIs: []string{
			"https://api.etherscan.io/api?module=account&action=balance&address=%s&tag=latest",
		},
		Symbol:   "ETH",
		Decimals: 18,
	},
	"bsc": {
		Name:            "Binance Smart Chain",
		DerivationPaths: []string{"m/44'/60'/0'/0/0"},
		APIs: []string{
			"https://api.bscscan.com/api?module=account&action=balance&address=%s",
		},
		Symbol:   "BNB",
		Decimals: 18,
	},
	"polygon": {
		Name:            "Polygon",
		DerivationPaths: []string{"m/44'/60'/0'/0/0"},
		APIs: []string{
			"https://api.polygonscan.com/api?module=account&action=balance&address=%s",
		},
		Symbol:   "MATIC",
		Decimals: 18,
	},
	"dogecoin": {
		Name:            "Dogecoin",
		DerivationPaths: []string{"m/44'/3'/0'/0/0"},
		APIs: []string{
			"https://dogechain.info/api/v1/address/balance/%s",
		},
		Symbol:   "DOGE",
		Decimals: 8,
	},
	"litecoin": {
		Name: "Litecoin",
		DerivationPaths: []string{
			"m/44'/2'/0'/0/0", // Legacy
			"m/84'/2'/0'/0/0", // Native SegWit
		},
		APIs: []string{
			"https://api.blockcypher.com/v1/ltc/main/addrs/%s/balance",
		},
		Symbol:   "LTC",
		Decimals: 8,
	},
}
