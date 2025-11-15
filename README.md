Blockchain Seed Checker 🔐
⚠️ SECURITY DISCLAIMER: This tool is for EDUCATIONAL AND RESEARCH purposes ONLY. Unauthorized use for accessing wallets you don't own is ILLEGAL.

A high-performance Go tool for educational analysis of BIP-39 mnemonic phrases and wallet address derivation across multiple blockchain networks.

🚀 Features
✅ BIP-39 Mnemonic Generator - Standard-compliant mnemonic generation

✅ HD Wallet Derivation - Support for BIP-32, BIP-44, BIP-84 paths

✅ Multi-Blockchain Support - Bitcoin, Ethereum, BSC, Polygon, Dogecoin, Litecoin

✅ Fast API-Based Scanning - Real-time balance checking

✅ Concurrent Execution - High-performance scanning

✅ Single Binary - Easy deployment with no dependencies

✅ Cross-Platform - Runs on Windows, macOS, and Linux

📦 Installation
Method 1: From Source
bash
# Clone repository
git clone https://github.com/RMotsId/blockchain-seed-checker.git
cd blockchain-seed-checker

# Build binary
go build -o seedchecker cmd/seedchecker/main.go

# Make executable
chmod +x seedchecker
Method 2: Go Install
bash
go install github.com/RMotsId/blockchain-seed-checker/cmd/seedchecker@latest
Method 3: Download Pre-built Binary
Check the Releases page for pre-built binaries.

🛠️ Usage
Check a single seed phrase:
bash
./seedchecker --seed "word1 word2 ... word12"
Generate and check random seeds:
bash
./seedchecker --generate 1000 --words 12 --batch-size 20
Check seeds from file:
bash
./seedchecker --file seeds.txt --passphrase "mypass"
Advanced options:
bash
./seedchecker --generate 5000 --words 24 --batch-size 30 --verbose --output-dir ./results
📋 Command Line Options
Option	Description	Default
--seed	Check a single seed phrase	-
--generate	Generate and check N seeds	0
--file	Check seeds from text file	-
--words	Words in mnemonic (12,15,18,21,24)	12
--passphrase	BIP-39 passphrase	""
--batch-size	Concurrent checking batch size	10
--output-dir	Custom output directory	current
--verbose	Enable verbose output	false
🏗️ Project Structure
text
blockchain-seed-checker/
├── cmd/seedchecker/          # CLI entry point
│   └── main.go              # Main application
├── internal/                 # Private application code
│   ├── generator/           # BIP-39 mnemonic generation
│   ├── deriver/             # HD wallet derivation  
│   ├── scanner/             # Blockchain scanning
│   └── utils/               # Utilities and logging
├── pkg/                     # Public library code
│   ├── config/              # Configuration
│   └── types/               # Shared types
├── go.mod                   # Go module definition
├── go.sum                   # Dependency checksums
├── LICENSE                  # MIT License
└── README.md               # This file
🔗 Supported Blockchains
Bitcoin (BTC) - Legacy, SegWit, Bech32

Ethereum (ETH) - ETH & ERC-20 tokens

Binance Smart Chain (BNB) - BEP-20 tokens

Polygon (MATIC)

Dogecoin (DOGE)

Litecoin (LTC)

🧪 Testing
bash
# Run tests
go test ./...

# Run tests with coverage
go test -cover ./...

# Build and test binary
go build -o seedchecker cmd/seedchecker/main.go
./seedchecker --help
📊 Example Output
bash
$ ./seedchecker --seed "abandon ability able about above absent absorb abstract absurd abuse access accident"

🔍 Checking single seed phrase...
📊 RESULTS:
============================================================
💰 bitcoin: 0.00000000 BTC
💰 ethereum: 0.00000000 ETH  
💰 bsc: 0.00000000 BNB
✅ No balances found
============================================================
⚠️ Legal Notice
IMPORTANT: This software is provided exclusively for security research, educational purposes, and authorized penetration testing.

❌ DO NOT use this tool to access wallets you do not own

❌ DO NOT use for illegal activities

✅ DO use for educational blockchain research

✅ DO use for testing your own wallets

✅ DO use for authorized security assessments

Users are solely responsible for complying with all applicable laws and regulations. The developers are not responsible for any misuse of this software.

🔒 Security Best Practices
🔐 Never use this tool with your own wallet seeds on untrusted systems

🛡️ Always run in isolated environments or virtual machines

📜 Use only for authorized security research and education

🔄 Regularly update dependencies for security patches

📝 Keep detailed logs of authorized usage

🐛 Reporting Issues
If you find any issues or have suggestions:

Check existing Issues

Create a new issue with detailed description

Include steps to reproduce if it's a bug

🤝 Contributing
We welcome contributions for educational improvements:

Fork the repository

Create a feature branch (git checkout -b feature/improvement)

Commit your changes (git commit -m 'Add some improvement')

Push to the branch (git push origin feature/improvement)

Open a Pull Request

📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

🙏 Acknowledgments
BIP-39 specification authors

Go Ethereum team

Bitcoin development community

All blockchain explorers providing public APIs

Remember: With great power comes great responsibility. Use this tool wisely and ethically. 🛡️
