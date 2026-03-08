# Assignment 01 — Simple Blockchain in Go

## 📌 Overview

This project implements a simplified blockchain in Go to demonstrate the core fundamentals of blockchain systems without Proof‑of‑Work. It focuses on how blocks are structured, how cryptographic hashes link blocks together, and how tampering can be detected through chain verification.

The implementation is designed for academic learning and demonstrates immutability, hash chaining, and integrity verification in a clear and minimal way.

---

## 🎯 Objectives

* Understand blockchain data structure
* Implement block creation and linking
* Generate cryptographic hashes using SHA‑256
* Ensure immutability through hash chaining
* Detect tampering by verifying the chain
* Demonstrate blockchain validation logic

---

## 🧱 Blockchain Features Implemented

### Block Structure

Each block contains:

* **Index** — Position of block in chain
* **Timestamp** — Creation time of block
* **Transaction** — Transaction data stored in block
* **Nonce** — Arbitrary number (no mining used)
* **PreviousHash** — Hash of previous block
* **Hash** — Current block hash

### Hash Calculation

* Uses **SHA‑256** cryptographic hashing
* Hash input concatenates:

  ```
  transaction + nonce + previousHash + index + timestamp
  ```
* Ensures data integrity and tamper evidence

### Hash Linking

Each block stores the previous block’s hash, forming a secure chain.

### Immutability

Any modification to block data changes its hash, breaking the chain.

### Tamper Detection

Chain verification logic detects:

* Invalid block hashes
* Broken previous hash links

### Not Included

* Proof‑of‑Work
* Mining
* Distributed networking
* Cryptocurrency logic

---

## 📁 Project Structure

```
assignment01bca_YourRollNo/
│
├── main.go
└── assignment01bca/
    └── blockchain.go
```

---

## ⚙️ Package Details

### Package Name

```
assignment01bca
```

### Core Functions

#### `NewBlock(transaction string, nonce int, previousHash string, index int) *Block`

Creates a new block and generates its hash.

#### `AddBlock(transaction string, nonce int)`

Adds a new block to the blockchain.

#### `ListBlocks()`

Prints all blocks with complete details.

#### `ChangeBlock(index int, newTransaction string)`

Modifies a block’s transaction (used to demonstrate tampering).

#### `VerifyChain() bool`

Validates blockchain integrity by checking:

* Recalculated hash vs stored hash
* Previous hash linkage

#### `CalculateHash(stringToHash string) string`

Generates SHA‑256 hash.

---

## 🎓 Custom Assignment Requirements

### 🔢 Roll Number Injection

* Last 3 digits of roll number included in at least one transaction

### 🧬 Personalized Genesis Block

* **Transaction:** `Genesis Block - YourRollNo`
* **Nonce:** Sum of digits of roll number

### 🧪 Tamper Demonstration

Program flow in `main.go`:

1. Create blockchain with Genesis block
2. Add at least 3 blocks
3. Print blockchain
4. Verify blockchain (valid)
5. Tamper with Block 1
6. Verify blockchain again (fails)

---

## ▶️ How to Run

### 1️⃣ Install Go

```
go version
```

### 2️⃣ Run Program

```
go run main.go
```

---

## 🖥️ Sample Output (Simplified)

```
========== BLOCKCHAIN ==========
Index: 0
Transaction: Genesis Block - FA21-BCS-123
Hash: 000abc...

Index: 1
Transaction: Alice pays Bob 10 BTC - 123
Hash: 91fa32...

Verifying blockchain...
Blockchain is valid

Tampering with Block 1...

Verifying blockchain...
Block 1 hash is invalid
```

---

## 🧠 Key Concepts Demonstrated

* Cryptographic hashing
* Data integrity
* Linked data structures
* Tamper detection
* Basic blockchain architecture

---

## 🧪 Testing Scenarios

| Scenario                  | Expected Result        |
| ------------------------- | ---------------------- |
| Normal chain verification | Valid blockchain       |
| Modify transaction data   | Hash mismatch detected |
| Break previous hash link  | Chain invalid          |

---

## 🏁 Submission Requirements

* GitHub repository named: `assignment01bca`
* ZIP file: `assignment01bca_YourRollNo.zip`
* Must contain:

  * `main.go`
  * `assignment01bca` package folder
* Live code demonstration required

---

## 👨‍💻 Author

**Name:** Muazam Ali
**Roll No:** 22i-1734
**Course:** Blockchain Fundamentals

---

## 📜 License

This project is for academic and educational use only.
