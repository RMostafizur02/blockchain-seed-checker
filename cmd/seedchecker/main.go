#!/bin/bash
# fix-all-imports.sh

# Clone the repository
git clone https://github.com/RMostafizur02/blockchain-seed-checker.git
cd blockchain-seed-checker

# Create fixed files with correct import paths

# 1. Fix go.mod
cat > go.mod << 'EOF'
module github.com/RMostafizur02/blockchain-seed-checker

go 1.21

require (
    github.com/btcsuite/btcd/btcutil v1.1.5
    github.com/ethereum/go-ethereum v1.13.5
    github.com/fatih/color v1.16.0
    github.com/tyler-smith/go-bip39 v1.1.0
    golang.org/x/crypto v0.17.0
)
EOF

# 2. Fix main.go with correct imports
cat > cmd/seedchecker/main.go << 'EOF'
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RMostafizur02/blockchain-seed-checker/internal/generator"
	"github.com/RMostafizur02/blockchain-seed-checker/internal/scanner"
	"github.com/RMostafizur02/blockchain-seed-checker/pkg/config"
)

var (
	generate    = flag.Int("generate", 0, "Generate and check N seed phrases")
	seed        = flag.String("seed", "", "Check a single seed phrase")
	file        = flag.String("file", "", "Check seed phrases from file")
	words       = flag.Int("words", 12, "Number of words in mnemonic (12,15,18,21,24)")
	passphrase  = flag.String("passphrase", "", "BIP-39 passphrase")
	batchSize   = flag.Int("batch-size", 10, "Batch size for concurrent checking")
	outputDir   = flag.String("output-dir", "", "Output directory for results")
	verbose     = flag.Bool("verbose", false, "Enable verbose output")
)

func printBanner() {
	banner := `
    ╔══════════════════════════════════════════════════════════════════════╗
    ║                BLOCKCHAIN SEED CHECKER (Go) v1.0.0                  ║
    ║                                                                    ║
    ║           🚨 EDUCATIONAL AND RESEARCH USE ONLY 🚨                 ║
    ║         Unauthorized access to wallets is ILLEGAL                 ║
    ║                                                                    ║
    ║      Supports: BTC, ETH, BSC, MATIC, DOGE, LTC + EVM chains       ║
    ╚══════════════════════════════════════════════════════════════════════╝
    `
	fmt.Println(banner)
}

func validateArgs() error {
	if *generate < 0 {
		return fmt.Errorf("number of seeds to generate must be positive")
	}
	if *batchSize <= 0 {
		return fmt.Errorf("batch size must be positive")
	}
	if *words != 12 && *words != 15 && *words != 18 && *words != 21 && *words != 24 {
		return fmt.Errorf("word count must be 12, 15, 18, 21, or 24")
	}
	if *file != "" {
		if _, err := os.Stat(*file); os.IsNotExist(err) {
			return fmt.Errorf("seed file not found: %s", *file)
		}
	}
	return nil
}

func main() {
	flag.Parse()
	
	printBanner()

	// Setup signal handling
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	
	go func() {
		<-sigChan
		fmt.Println("\n⏹️  Operation cancelled by user")
		os.Exit(130)
	}()

	// Validate arguments
	if err := validateArgs(); err != nil {
		log.Fatalf("❌ Argument error: %v", err)
	}

	// Create output directory if specified
	if *outputDir != "" {
		if err := os.MkdirAll(*outputDir, 0755); err != nil {
			log.Fatalf("❌ Failed to create output directory: %v", err)
		}
	}

	startTime := time.Now()
	
	var err error
	switch {
	case *seed != "":
		err = checkSingleSeed()
	case *file != "":
		err = checkFromFile()
	case *generate > 0:
		err = generateAndCheck()
	default:
		flag.Usage()
		return
	}

	if err != nil {
		log.Fatalf("❌ Error: %v", err)
	}

	elapsed := time.Since(startTime)
	fmt.Printf("⏱️  Total execution time: %v\n", elapsed)
}

func checkSingleSeed() error {
	fmt.Printf("🔍 Checking single seed phrase...\n")
	
	// Validate mnemonic
	mnemonic := *seed
	if !generator.ValidateMnemonic(mnemonic) {
		return fmt.Errorf("invalid BIP-39 mnemonic phrase")
	}

	// Generate seed from mnemonic
	seed, err := generator.MnemonicToSeed(mnemonic, *passphrase)
	if err != nil {
		return fmt.Errorf("failed to generate seed: %v", err)
	}

	// Derive addresses
	addresses, err := scanner.DeriveAllAddresses(seed)
	if err != nil {
		return fmt.Errorf("failed to derive addresses: %v", err)
	}

	// Scan addresses
	results, err := scanner.ScanAddresses(addresses)
	if err != nil {
		return fmt.Errorf("failed to scan addresses: %v", err)
	}

	// Print results
	printResults(mnemonic, results)
	return nil
}

func checkFromFile() error {
	fmt.Printf("📁 Checking seeds from file: %s\n", *file)
	
	content, err := os.ReadFile(*file)
	if err != nil {
		return fmt.Errorf("failed to read file: %v", err)
	}

	// TODO: Implement file processing
	fmt.Println("File processing not yet implemented")
	return nil
}

func generateAndCheck() error {
	fmt.Printf("🎯 Generating and checking %d seed phrases...\n", *generate)
	
	// TODO: Implement batch generation and checking
	fmt.Println("Batch generation not yet implemented")
	return nil
}

func printResults(mnemonic string, results map[string]config.ScanResult) {
	fmt.Printf("\n📊 RESULTS for: %s\n", mnemonic)
	fmt.Println("=" + string(make([]byte, 60)) + "=")
	
	totalBalance := 0.0
	hasBalance := false
	
	for chain, result := range results {
		if result.Balance > 0 {
			hasBalance = true
			totalBalance += result.Balance
			fmt.Printf("💰 %s: %.8f %s\n", chain, result.Balance, result.Unit)
			fmt.Printf("   Address: %s\n", result.Address)
		}
	}
	
	if hasBalance {
		fmt.Printf("\n💵 TOTAL BALANCE: %.8f\n", totalBalance)
		fmt.Printf("🔍 Check log file for detailed results\n")
	} else {
		fmt.Println("❌ No balances found")
	}
	
	fmt.Println("=" + string(make([]byte, 60)) + "=")
}
EOF

# 3. Fix generator/bip39.go
cat > internal/generator/bip39.go << 'EOF'
package generator

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"

	"github.com/RMostafizur02/blockchain-seed-checker/pkg/config"
	"github.com/tyler-smith/go-bip39"
)

// GenerateMnemonic generates a BIP-39 compliant mnemonic phrase
func GenerateMnemonic(wordCount int) (string, error) {
	// Map word count to entropy bits
	entropyBits := map[int]int{
		12: 128,
		15: 160,
		18: 192,
		21: 224,
		24: 256,
	}

	bits, ok := entropyBits[wordCount]
	if !ok {
		return "", fmt.Errorf("invalid word count: %d. Must be 12, 15, 18, 21, or 24", wordCount)
	}

	entropy := make([]byte, bits/8)
	_, err := rand.Read(entropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate entropy: %v", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy)
	if err != nil {
		return "", fmt.Errorf("failed to generate mnemonic: %v", err)
	}

	return mnemonic, nil
}

// ValidateMnemonic validates a BIP-39 mnemonic phrase
func ValidateMnemonic(mnemonic string) bool {
	return bip39.IsMnemonicValid(mnemonic)
}

// MnemonicToSeed converts a mnemonic to a seed using BIP-39 specification
func MnemonicToSeed(mnemonic, passphrase string) ([]byte, error) {
	if !ValidateMnemonic(mnemonic) {
		return nil, fmt.Errorf("invalid mnemonic phrase")
	}

	seed := bip39.NewSeed(mnemonic, passphrase)
	return seed, nil
}

// GenerateMnemonicCustom generates a mnemonic using custom wordlist (for educational purposes)
func GenerateMnemonicCustom(wordCount int) (string, error) {
	wordlist := config.GetWordlist()
	
	entropyBits := map[int]int{
		12: 128,
		15: 160,
		18: 192,
		21: 224,
		24: 256,
	}

	bits, ok := entropyBits[wordCount]
	if !ok {
		return "", fmt.Errorf("invalid word count: %d", wordCount)
	}

	// Generate entropy
	entropy := make([]byte, bits/8)
	_, err := rand.Read(entropy)
	if err != nil {
		return "", err
	}

	// Calculate checksum
	hash := sha256.Sum256(entropy)
	checksumBits := bits / 32
	checksumByte := hash[0] >> (8 - uint(checksumBits))

	// Convert entropy to binary string
	var entropyBitsStr string
	for _, b := range entropy {
		entropyBitsStr += fmt.Sprintf("%08b", b)
	}

	// Add checksum
	entropyBitsStr += fmt.Sprintf("%0*b", checksumBits, checksumByte)

	// Split into 11-bit chunks and map to words
	var words []string
	for i := 0; i < len(entropyBitsStr); i += 11 {
		end := i + 11
		if end > len(entropyBitsStr) {
			end = len(entropyBitsStr)
		}
		
		chunk := entropyBitsStr[i:end]
		index := binaryFromString(chunk)
		if index >= len(wordlist) {
			return "", fmt.Errorf("word index out of range")
		}
		words = append(words, wordlist[index])
	}

	return strings.Join(words, " "), nil
}

func binaryFromString(s string) uint16 {
	var result uint16
	for _, ch := range s {
		result = result<<1 | uint16(ch-'0')
	}
	return result
}
EOF

# Update dependencies
go mod tidy

# Test the build
echo "🔨 Testing build..."
go build -o seedchecker cmd/seedchecker/main.go

if [ $? -eq 0 ]; then
    echo "✅ Build successful! All imports fixed."
    echo "🚀 Committing changes..."
    
    git add .
    git commit -m "Fix all import paths to match repository name"
    git push origin main
    
    echo "✅ All files updated successfully!"
    echo "📦 Now install with: go install github.com/RMostafizur02/blockchain-seed-checker/cmd/seedchecker@latest"
else
    echo "❌ Build failed. Please check the errors above."
fi
