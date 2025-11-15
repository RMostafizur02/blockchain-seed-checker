# Crypto Seed Scanner 🔐

> ⚠️ **SECURITY DISCLAIMER**: This tool is for **EDUCATIONAL AND RESEARCH purposes ONLY**. Unauthorized use for accessing wallets you don't own is **ILLEGAL**.

A high-performance Go tool for analyzing BIP-39 mnemonic phrases and deriving wallet addresses across multiple blockchain networks.

## 🚀 Features

- ✅ **BIP-39 Mnemonic Generator** - Standard-compliant mnemonic generation
- ✅ **HD Wallet Derivation** - Support for BIP-32, BIP-44, BIP-84 paths  
- ✅ **Multi-Blockchain Support** - Bitcoin, Ethereum, BSC, Polygon, Dogecoin, Litecoin
- ✅ **Fast API-Based Scanning** - Real-time balance checking
- ✅ **Concurrent Execution** - High-performance scanning
- ✅ **Single Binary** - Easy deployment with no dependencies

## 📦 Installation

```bash
# Clone repository
git clone https://github.com/RMotsId/crypto-seed-scanner.git
cd crypto-seed-scanner

# Build binary
go build -o seedscanner cmd/seedchecker/main.go

# Or install directly
go install github.com/RMotsId/crypto-seed-scanner/cmd/seedchecker@latest
🛠️ Usage
Check a single seed phrase:
bash
./seedscanner --seed "word1 word2 ... word12"
Generate and check random seeds:
bash
./seedscanner --generate 1000 --words 12 --batch-size 20
Check seeds from file:
bash
./seedscanner --file seeds.txt --passphrase "mypass"
Advanced options:
bash
./seedscanner --generate 5000 --words 24 --batch-size 30 --verbose --output-dir ./results
🏗️ Project Structure
text
crypto-seed-scanner/
├── cmd/seedchecker/     # CLI entry point
├── internal/            # Private application code
│   ├── generator/       # BIP-39 mnemonic generation
│   ├── deriver/         # HD wallet derivation  
│   ├── scanner/         # Blockchain scanning
│   └── utils/           # Utilities and logging
├── pkg/                 # Public library code
│   ├── config/          # Configuration
│   └── types/           # Shared types
└── go.mod              # Go module definition
⚠️ Legal Notice
This software is provided for security research and educational purposes only. Users are responsible for complying with all applicable laws and regulations. The developers are not responsible for any misuse of this software.

🔒 Security Best Practices
Never use this tool with your own wallet seeds

Always run in isolated environments

Use only for authorized security research

Regularly update dependencies

📄 License
MIT License - see LICENSE file for details
