package types

type WalletAddress struct {
	Chain     string
	Address   string
	PublicKey string
	Path      string
}

type ScanResult struct {
	Chain       string
	Address     string
	Balance     float64
	Unit        string
	HasBalance  bool
	Error       string
}

type SeedCheckResult struct {
	Mnemonic    string
	Passphrase  string
	Addresses   map[string][]WalletAddress
	ScanResults map[string]ScanResult
	HasBalance  bool
	TotalBalance float64
	Error       string
	Timestamp   int64
}
