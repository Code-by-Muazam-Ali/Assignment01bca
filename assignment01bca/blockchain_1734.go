// ===============================
// File: assignment01bca/blockchain.go
// Package: assignment01bca
// ===============================

package assignment01bca

import (
    "crypto/sha256"
    "encoding/hex"
    "fmt"
    "strconv"
    "time"
)

type Block struct {
    Index        int
    Timestamp    string
    Transaction  string
    Nonce        int
    PreviousHash string
    Hash         string
}

type Blockchain struct {
    Blocks []*Block
}

// -------------------------------
// CalculateHash
// -------------------------------
func CalculateHash(stringToHash string) string {
    h := sha256.Sum256([]byte(stringToHash))
    return hex.EncodeToString(h[:])
}

// -------------------------------
// CreateHash (internal helper)
// -------------------------------
func (b *Block) CreateHash() string {
    record := b.Transaction +
        strconv.Itoa(b.Nonce) +
        b.PreviousHash +
        strconv.Itoa(b.Index) +
        b.Timestamp

    return CalculateHash(record)
}

// -------------------------------
// NewBlock
// -------------------------------
func NewBlock(transaction string, nonce int, previousHash string, index int) *Block {
    block := &Block{
        Index:        index,
        Timestamp:    time.Now().Format(time.RFC3339),
        Transaction:  transaction,
        Nonce:        nonce,
        PreviousHash: previousHash,
    }
    block.Hash = block.CreateHash()
    return block
}

// -------------------------------
// CreateBlockchain
// -------------------------------
func CreateBlockchain(rollNo string) *Blockchain {
    bc := &Blockchain{}

    // Sum of digits of roll number
    sum := 0
    for _, ch := range rollNo {
        if ch >= '0' && ch <= '9' {
            sum += int(ch - '0')
        }
    }

    genesisTx := "Genesis Block - " + rollNo
    genesis := NewBlock(genesisTx, sum, "0", 0)

    bc.Blocks = append(bc.Blocks, genesis)
    return bc
}

// -------------------------------
// AddBlock
// -------------------------------
func (bc *Blockchain) AddBlock(transaction string, nonce int) {
    prev := bc.Blocks[len(bc.Blocks)-1]
    newBlock := NewBlock(transaction, nonce, prev.Hash, prev.Index+1)
    bc.Blocks = append(bc.Blocks, newBlock)
}

// -------------------------------
// ListBlocks
// -------------------------------
func (bc *Blockchain) ListBlocks() {
    fmt.Println("\n========== BLOCKCHAIN ==========")
    for _, b := range bc.Blocks {
        fmt.Println("--------------------------------")
        fmt.Println("Index:        ", b.Index)
        fmt.Println("Timestamp:    ", b.Timestamp)
        fmt.Println("Transaction:  ", b.Transaction)
        fmt.Println("Nonce:        ", b.Nonce)
        fmt.Println("Previous Hash:", b.PreviousHash)
        fmt.Println("Hash:         ", b.Hash)
    }
    fmt.Println("--------------------------------")
}

// -------------------------------
// ChangeBlock (tamper)
// -------------------------------
func (bc *Blockchain) ChangeBlock(index int, newTransaction string) {
    if index < 0 || index >= len(bc.Blocks) {
        fmt.Println("Invalid block index")
        return
    }
    bc.Blocks[index].Transaction = newTransaction
    bc.Blocks[index].Hash = bc.Blocks[index].CreateHash()
}

// -------------------------------
// VerifyChain
// -------------------------------
func (bc *Blockchain) VerifyChain() bool {
    fmt.Println("\nVerifying blockchain...")

    for i := 1; i < len(bc.Blocks); i++ {
        curr := bc.Blocks[i]
        prev := bc.Blocks[i-1]

        // Recalculate hash
        if curr.Hash != curr.CreateHash() {
            fmt.Println(" Block", curr.Index, "hash is invalid")
            return false
        }

        // Check link
        if curr.PreviousHash != prev.Hash {
            fmt.Println(" Block", curr.Index, "previous hash mismatch")
            return false
        }
    }

    fmt.Println(" Blockchain is valid")
    return true
}
