package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"crypto-wallet-seed-checker/internal/generator"
	"crypto-wallet-seed-checker/internal/scanner"
	"crypto-wallet-seed-checker/pkg/config"
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
    ║                CRYPTO WALLET SEED CHECKER (Go) v1.0.0               ║
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
