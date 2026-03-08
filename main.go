// ===============================
// File: main.go
// ===============================

package main

import (
    "fmt"
    "assignment01bca"
)

func main() {
    rollNo := "22i-1734"     
    last3 := "734"           

    bc := assignment01bca.CreateBlockchain(rollNo)

    // Create 10 blocks
    bc.AddBlock("Ali pays Hamza 10 BTC - "+last3, 101)
    bc.AddBlock("Hamza pays Asim 5 BTC", 102)
    bc.AddBlock("Asim pays Umer 3 BTC", 103)
    bc.AddBlock("Umer pays Mirza 7 BTC", 104)
    bc.AddBlock("Mirza pays Mustansir 2 BTC", 105)
    bc.AddBlock("Mustansir pays Bilal 6 BTC", 106)
    bc.AddBlock("Bilal pays Ahmed 4 BTC", 107)
    bc.AddBlock("Ahmed pays Saif 8 BTC", 108)
    bc.AddBlock("Saif pays Ali 1 BTC", 109)
    bc.AddBlock("Ali pays Umer 9 BTC", 110)

    // Print blockchain
    bc.ListBlocks()

    // Verify blockchain
    bc.VerifyChain()

    // Tamper with Block 1
    fmt.Println("\n Tampering with Block 1...")
    bc.ChangeBlock(1, "Alice pays Mallory 1000 BTC - HACKED")

    // Verify again
    bc.VerifyChain()
}
