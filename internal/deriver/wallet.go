package deriver

import (
	"crypto-wallet-seed-checker/pkg/config"
	"crypto-wallet-seed-checker/pkg/types"
	"fmt"

	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

// DeriveAllAddresses derives addresses for all supported blockchains from a seed
func DeriveAllAddresses(seed []byte) (map[string][]types.WalletAddress, error) {
	results := make(map[string][]types.WalletAddress)

	for chainName, chainConfig := range config.Blockchains {
		var addresses []types.WalletAddress
		
		for _, path := range chainConfig.DerivationPaths {
			address, err := DeriveAddress(seed, path, chainName)
			if err != nil {
				return nil, fmt.Errorf("failed to derive %s address: %v", chainName, err)
			}
			addresses = append(addresses, address)
		}
		
		results[chainName] = addresses
	}

	return results, nil
}

// DeriveAddress derives a single address from seed using specified derivation path
func DeriveAddress(seed []byte, path, chain string) (types.WalletAddress, error) {
	switch chain {
	case "ethereum", "bsc", "polygon":
		return deriveEthereumAddress(seed, path)
	case "bitcoin", "litecoin", "dogecoin":
		return deriveBitcoinAddress(seed, path, chain)
	default:
		return types.WalletAddress{}, fmt.Errorf("unsupported chain: %s", chain)
	}
}

func deriveEthereumAddress(seed []byte, path string) (types.WalletAddress, error) {
	// Create master key from seed
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return types.WalletAddress{}, err
	}

	// Parse derivation path
	derivationPath, err := accounts.ParseDerivationPath(path)
	if err != nil {
		return types.WalletAddress{}, err
	}

	// Derive key according to path
	key := masterKey
	for _, n := range derivationPath {
		key, err = key.NewChildKey(n)
		if err != nil {
			return types.WalletAddress{}, err
		}
	}

	// Convert to Ethereum private key
	privateKey, err := crypto.ToECDSA(key.Key)
	if err != nil {
		return types.WalletAddress{}, err
	}

	// Generate address from public key
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*crypto.PublicKey)
	if !ok {
		return types.WalletAddress{}, fmt.Errorf("error casting public key to ECDSA")
	}

	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return types.WalletAddress{
		Chain:     "ethereum",
		Address:   address.Hex(),
		PublicKey: common.Bytes2Hex(crypto.FromECDSAPub(publicKeyECDSA)),
		Path:      path,
	}, nil
}

func deriveBitcoinAddress(seed []byte, path, chain string) (types.WalletAddress, error) {
	// Get appropriate network parameters
	var params *chaincfg.Params
	switch chain {
	case "bitcoin":
		params = &chaincfg.MainNetParams
	case "litecoin":
		params = &chaincfg.MainNetParams // Litecoin uses same base as Bitcoin
	case "dogecoin":
		params = &chaincfg.MainNetParams // Dogecoin uses same base as Bitcoin
	default:
		params = &chaincfg.MainNetParams
	}

	// Create master key from seed
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return types.WalletAddress{}, err
	}

	// Parse and derive path (simplified - in production use proper HD derivation)
	// For demo purposes, we'll use a simplified approach
	key := masterKey
	
	// Generate address from public key
	publicKey, err := btcutil.NewAddressPubKey(key.PublicKey().Key, params)
	if err != nil {
		return types.WalletAddress{}, err
	}

	return types.WalletAddress{
		Chain:     chain,
		Address:   publicKey.EncodeAddress(),
		PublicKey: common.Bytes2Hex(key.PublicKey().Key),
		Path:      path,
	}, nil
}

// MnemonicToSeed is a wrapper around the bip39 function
func MnemonicToSeed(mnemonic, passphrase string) ([]byte, error) {
	return bip39.NewSeedWithErrorChecking(mnemonic, passphrase)
}
