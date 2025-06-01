# Post-Quantum Cryptography POC

This is a minimal proof-of-concept implementation of **Post-Quantum Cryptography (PQC)**  

## What This Demonstrates

This POC implements:
- **Post-Quantum Cryptography (PQC) Integration**: Using NIST-approved lattice-based cryptography (Dilithium)
- **Future-proof smart contracts**: Quantum-resistant digital signatures for transactions

## Technical Details

- **Algorithm**: Dilithium Mode 3 (lattice-based signatures)
- **Security Level**: Equivalent to AES-192
- **Standard**: NIST Post-Quantum Cryptography approved

## Running the Demo

```bash
# Install dependencies
go mod tidy

# Run the POC
go run main.go
```

## Sample Output

```
🚀 Post-Quantum Cryptography POC
============================================================
✅ Generated keypair using Dilithium-Mode3
   Public Key Size: 1952 bytes
   Private Key Size: 4000 bytes

📝 Creating quantum-resistant transaction:
   From: quantum_wallet_1
   To: quantum_wallet_2
   Amount: 1000 COINS
   Data: quantum-secured transfer

🔐 Transaction signed with quantum-resistant signature
   Signature Size: 3293 bytes
   Signature (first 32 bytes): a1b2c3d4...

✅ Signature verification PASSED
   Transaction is quantum-secure and valid!

🛡️  This demonstrates QuantumLayer's core PQC capability:
   - Lattice-based cryptography (Dilithium)
   - Quantum-resistant digital signatures
   - Future-proof transaction security
   - Ready for post-quantum era
```
