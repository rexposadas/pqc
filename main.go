package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"strings"

	"github.com/cloudflare/circl/sign/dilithium"
)

type QuantumSecureWallet struct {
	PublicKey  []byte
	PrivateKey []byte
	Algorithm  string
}

func NewQuantumSecureWallet() (*QuantumSecureWallet, error) {
	// Using Dilithium (NIST-approved lattice-based signature scheme)
	mode := dilithium.Mode3 // Security level equivalent to AES-192

	publicKey, privateKey, err := mode.GenerateKey(rand.Reader)
	if err != nil {
		return nil, fmt.Errorf("failed to generate quantum-secure keys: %w", err)
	}

	return &QuantumSecureWallet{
		PublicKey:  publicKey.Bytes(),
		PrivateKey: privateKey.Bytes(),
		Algorithm:  "Dilithium-Mode3",
	}, nil
}

func (w *QuantumSecureWallet) SignTransaction(transaction []byte) ([]byte, error) {
	mode := dilithium.Mode3
	privateKey := mode.PrivateKeyFromBytes(w.PrivateKey)

	signature := mode.Sign(privateKey, transaction)
	return signature, nil
}

func (w *QuantumSecureWallet) VerifySignature(transaction []byte, signature []byte) bool {
	mode := dilithium.Mode3
	publicKey := mode.PublicKeyFromBytes(w.PublicKey)

	return mode.Verify(publicKey, transaction, signature)
}

// Simulate a simple quantum-resistant smart contract transaction
type Transaction struct {
	From   string
	To     string
	Amount uint64
	Data   []byte
}

func (t *Transaction) Serialize() []byte {
	return []byte(fmt.Sprintf("%s:%s:%d:%x", t.From, t.To, t.Amount, t.Data))
}

func main() {
	fmt.Println("🚀 Post-Quantum Cryptography POC")
	fmt.Println(strings.Repeat("=", 60))

	// Create quantum-secure wallet
	wallet, err := NewQuantumSecureWallet()
	if err != nil {
		log.Fatal("Failed to create quantum-secure wallet:", err)
	}

	fmt.Printf("✅ Generated quantum-resistant keypair using %s\n", wallet.Algorithm)
	fmt.Printf("   Public Key Size: %d bytes\n", len(wallet.PublicKey))
	fmt.Printf("   Private Key Size: %d bytes\n", len(wallet.PrivateKey))
	fmt.Println()

	// Create a sample transaction
	tx := &Transaction{
		From:   "quantum_wallet_1",
		To:     "quantum_wallet_2",
		Amount: 1000,
		Data:   []byte("quantum-secured transfer"),
	}

	fmt.Println("📝 Creating quantum-resistant transaction:")
	fmt.Printf("   From: %s\n", tx.From)
	fmt.Printf("   To: %s\n", tx.To)
	fmt.Printf("   Amount: %d QNTM\n", tx.Amount)
	fmt.Printf("   Data: %s\n", string(tx.Data))
	fmt.Println()

	// Sign the transaction with post-quantum cryptography
	txData := tx.Serialize()
	signature, err := wallet.SignTransaction(txData)
	if err != nil {
		log.Fatal("Failed to sign transaction:", err)
	}

	fmt.Printf("🔐 Transaction signed with quantum-resistant signature\n")
	fmt.Printf("   Signature Size: %d bytes\n", len(signature))
	fmt.Printf("   Signature (first 32 bytes): %x...\n", signature[:32])
	fmt.Println()

	// Verify the signature
	isValid := wallet.VerifySignature(txData, signature)
	if isValid {
		fmt.Println("✅ Signature verification PASSED")
		fmt.Println("   Transaction is quantum-secure and valid!")
	} else {
		fmt.Println("❌ Signature verification FAILED")
	}

	fmt.Println()
	fmt.Println("🛡️  This demonstrates PQC:")
	fmt.Println("   - Lattice-based cryptography (Dilithium)")
	fmt.Println("   - Future-proof transaction security")
}
