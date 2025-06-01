package main

import (
    "bytes"
    "testing"
)

func TestNewSecureWallet(t *testing.T) {
    wallet, err := NewSecureWallet()
    if err != nil {
        t.Fatalf("Failed to create wallet: %v", err)
    }
    
    if len(wallet.PublicKey) == 0 {
        t.Error("Public key should not be empty")
    }
    
    if len(wallet.PrivateKey) == 0 {
        t.Error("Private key should not be empty")
    }
    
    if wallet.Algorithm != "Dilithium-Mode3" {
        t.Errorf("Expected algorithm Dilithium-Mode3, got %s", wallet.Algorithm)
    }
}

func TestSignAndVerify(t *testing.T) {
    wallet, err := NewSecureWallet()
    if err != nil {
        t.Fatalf("Failed to create wallet: %v", err)
    }
    
    // Test data
    tx := &Transaction{
        From:   "wallet1",
        To:     "wallet2",
        Amount: 500,
        Data:   []byte("test transaction"),
    }
    
    txData := tx.Serialize()
    
    // Sign transaction
    signature, err := wallet.SignTransaction(txData)
    if err != nil {
        t.Fatalf("Failed to sign transaction: %v", err)
    }
    
    // Verify valid signature
    if !wallet.VerifySignature(txData, signature) {
        t.Error("Signature verification failed for valid signature")
    }
    
    // Verify tampered data fails
    tamperedData := bytes.Clone(txData)
    tamperedData[0] = tamperedData[0] + 1
    if wallet.VerifySignature(tamperedData, signature) {
        t.Error("Signature verification should fail for tampered data")
    }
}

func TestTransaction_Serialize(t *testing.T) {
    tx := &Transaction{
        From:   "wallet1",
        To:     "wallet2",
        Amount: 100,
        Data:   []byte("test"),
    }
    
    serialized := tx.Serialize()
    expected := "wallet1:wallet2:100:74657374"
    
    if string(serialized) != expected {
        t.Errorf("Expected serialized data %s, got %s", expected, string(serialized))
    }
}