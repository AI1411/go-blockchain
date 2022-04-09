package main

import (
	"blockchain/model"
	"fmt"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	myBlockchainAddress := "myBlockchainAddress"
	blockchain := model.NewBlockchain(myBlockchainAddress)
	blockchain.Print()

	blockchain.AddTransaction("A", "B", 5)
	blockchain.Mining()
	blockchain.Print()

	blockchain.AddTransaction("C", "D", 10)
	blockchain.AddTransaction("X", "Y", 15)
	blockchain.Mining()
	blockchain.Print()

	fmt.Printf("my %.1f\n", blockchain.CalculateTotalAmount("myBlockchainAddress"))
	fmt.Printf("C %.1f\n", blockchain.CalculateTotalAmount("C"))
	fmt.Printf("D %.1f\n", blockchain.CalculateTotalAmount("D"))
}
