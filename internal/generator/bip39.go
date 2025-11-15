package generator

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"strings"

	"crypto-wallet-seed-checker/pkg/config"
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
