🔐 Blockchain Seed Checker
⚠️ SECURITY DISCLAIMER: This tool is for EDUCATIONAL AND RESEARCH purposes ONLY. Unauthorized use for accessing wallets you don't own is ILLEGAL.

https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go
https://img.shields.io/badge/License-MIT-green?style=for-the-badge
https://img.shields.io/badge/Platform-Windows%2520%257C%2520macOS%2520%257C%2520Linux-lightgrey?style=for-the-badge

A high-performance Go tool for educational analysis of BIP-39 mnemonic phrases and wallet address derivation across multiple blockchain networks.

✨ Features
Feature	Description
🔑 BIP-39 Generator	Standard-compliant mnemonic generation
🗝️ HD Wallet Derivation	Support for BIP-32, BIP-44, BIP-84 paths
⛓️ Multi-Blockchain	Bitcoin, Ethereum, BSC, Polygon, Dogecoin, Litecoin
🚀 Fast API Scanning	Real-time balance checking
⚡ Concurrent Execution	High-performance parallel scanning
📦 Single Binary	Easy deployment with no dependencies
🌐 Cross-Platform	Runs on Windows, macOS, and Linux
🚀 Quick Start
Installation
bash
# Clone the repository
git clone https://github.com/RMotsId/blockchain-seed-checker.git
cd blockchain-seed-checker

# Build the binary
go build -o seedchecker cmd/seedchecker/main.go

# Make executable
chmod +x seedchecker
Or install directly:
bash
go install github.com/RMotsId/blockchain-seed-checker/cmd/seedchecker@latest
📖 Usage Examples
🔍 Check a Single Seed
bash
./seedchecker --seed "abandon ability able about above absent absorb abstract absurd abuse access accident"
🎯 Generate & Check Multiple Seeds
bash
./seedchecker --generate 1000 --words 12 --batch-size 20
📁 Check Seeds from File
bash
./seedchecker --file seeds.txt --passphrase "mypass"
⚡ Advanced Usage
bash
./seedchecker --generate 5000 --words 24 --batch-size 30 --verbose --output-dir ./results
🛠️ Command Reference
Command	Description	Default
--seed	Check single seed phrase	-
--generate	Generate & check N seeds	0
--file	Check seeds from file	-
--words	Words in mnemonic	12
--passphrase	BIP-39 passphrase	""
--batch-size	Concurrent batch size	10
--output-dir	Output directory	current
--verbose	Enable verbose output	false
🏗️ Project Structure
text
blockchain-seed-checker/
├── 📁 cmd/seedchecker/
│   └── 🎯 main.go                 # CLI entry point
├── 📁 internal/
│   ├── 🔑 generator/              # BIP-39 generation
│   ├── 🗝️ deriver/               # HD wallet derivation
│   ├── 🔍 scanner/               # Blockchain scanning
│   └── 🛠️ utils/                 # Utilities & logging
├── 📁 pkg/
│   ├── ⚙️ config/                # Configuration
│   └── 📊 types/                 # Shared types
├── 📄 go.mod                     # Dependencies
├── 📄 LICENSE                    # MIT License
└── 📄 README.md                 # This file
⛓️ Supported Blockchains
Blockchain	Support	APIs
Bitcoin (BTC)	✅ Legacy, SegWit, Bech32	Blockstream, Blockchain.com
Ethereum (ETH)	✅ ETH & ERC-20	Etherscan
Binance Chain (BNB)	✅ BEP-20 tokens	BscScan
Polygon (MATIC)	✅	Polygonscan
Dogecoin (DOGE)	✅	Dogechain
Litecoin (LTC)	✅	BlockCypher
📊 Example Output
bash
$ ./seedchecker --seed "your seed phrase here"

🎯 Blockchain Seed Checker v1.0.0
===========================================

🔍 Checking: your seed phrase here...
⏱️  Deriving addresses across 6 blockchains...

📊 SCAN RESULTS:
===========================================
✅ Bitcoin:   0.00000000 BTC
✅ Ethereum:  0.00000000 ETH
✅ BSC:       0.00000000 BNB
✅ Polygon:   0.00000000 MATIC
✅ Dogecoin:  0.00000000 DOGE
✅ Litecoin:  0.00000000 LTC
===========================================

💡 No balances found across all networks
⏱️  Scan completed in 2.3 seconds
🧪 Testing
bash
# Run all tests
go test ./...

# Run with coverage
go test -cover ./...

# Test specific package
go test ./internal/generator
⚠️ Legal & Security
🚫 Prohibited Uses
❌ Accessing wallets you don't own

❌ Illegal activities

❌ Unauthorized penetration testing

✅ Approved Uses
✅ Educational research

✅ Security coursework

✅ Authorized testing

✅ Personal wallet recovery

Warning: Users are solely responsible for legal compliance. Developers assume no liability for misuse.

🤝 Contributing
We welcome educational improvements:

🍴 Fork the repository

🌿 Create a feature branch: git checkout -b feature/improvement

💾 Commit changes: git commit -m 'Add educational feature'

📤 Push to branch: git push origin feature/improvement

🔄 Open a Pull Request

📄 License
This project is licensed under the MIT License - see the LICENSE file for details.

🙏 Acknowledgments
BIP-39 Specification Authors

Go Ethereum Team

Bitcoin Development Community

Blockchain Explorer API Providers

<div align="center">
🔐 Use Responsibly • 🛡️ Stay Legal • 📚 Learn Ethically

With great power comes great responsibility

</div>
📞 Support
🐛 Report Issues

💡 Request Features

📚 Read Documentation

<div align="center">
Made with ❤️ for the blockchain education community

</div>
This README now features:

🎨 Professional formatting with tables and icons

🛡️ Clear security warnings

📱 Mobile-responsive design

🚀 Quick start section

📖 Comprehensive examples

⚡ Visual command reference

🔗 Badges for professionalism

📊 Structured information
