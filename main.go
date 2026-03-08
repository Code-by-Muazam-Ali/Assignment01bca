package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"assignment01bca"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Cyan   = "\033[36m"
	Bold   = "\033[1m"
)

func main() {
	rollNo := "22i-1734" 
	last3 := "734"           

	bc := assignment01bca.CreateBlockchain(rollNo)
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println(Bold + Cyan + "\n===== Simple Blockchain CLI =====" + Reset)
		fmt.Println(Blue + "1." + Reset + " Add Block")
		fmt.Println(Blue + "2." + Reset + " List Blocks")
		fmt.Println(Blue + "3." + Reset + " Verify Blockchain")
		fmt.Println(Blue + "4." + Reset + " Tamper Block")
		fmt.Println(Blue + "5." + Reset + " Demo: Add 10 Sample Blocks")
		fmt.Println(Blue + "0." + Reset + " Exit")
		fmt.Print(Yellow + "Choose option: " + Reset)

		input, _ := reader.ReadString('\n')
		choice := strings.TrimSpace(input)

		switch choice {
		case "1":
			fmt.Print("Enter transaction (like 'Alice pays Bob 10 BTC'): ")
			tx, _ := reader.ReadString('\n')
			tx = strings.TrimSpace(tx)

			fmt.Print("Enter nonce (number, like 101): ")
			nStr, _ := reader.ReadString('\n')
			nStr = strings.TrimSpace(nStr)
			nonce, err := strconv.Atoi(nStr)
			if err != nil {
				fmt.Println(Red + "Invalid nonce" + Reset)
				continue
			}

			bc.AddBlock(tx, nonce)
			fmt.Println(Green + " Block added" + Reset)

		case "2":
			bc.ListBlocks()

		case "3":
			bc.VerifyChain()

		case "4":
			fmt.Print("Enter block index to tamper (like 1): ")
			iStr, _ := reader.ReadString('\n')
			iStr = strings.TrimSpace(iStr)
			idx, err := strconv.Atoi(iStr)
			if err != nil {
				fmt.Println(Red + "Invalid index" + Reset)
				continue
			}

			fmt.Print("Enter new transaction (like 'Bob pays Alice 5 BTC'): ")
			tx, _ := reader.ReadString('\n')
			tx = strings.TrimSpace(tx)

			bc.ChangeBlock(idx, tx)
			fmt.Println(Yellow + " Block modified" + Reset)

		case "5":
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
			fmt.Println(Green + " 10 sample blocks added" + Reset)

		case "0":
			fmt.Println("Goodbye :)")
			return

		default:
			fmt.Println(Red + "Invalid option" + Reset)
		}
	}
}
