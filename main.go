package main

import (
	"blockchain/model"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	myBlockchainAddress := "myBlockchainAddress"
	blockchain := model.NewBlockchain(myBlockchainAddress)
	blockchain.Print()

	blockchain.AddTransaction("Alice", "Bob", 5)
	blockchain.Mining()
	blockchain.Print()

	blockchain.AddTransaction("Bob", "Alice", 10)
	blockchain.AddTransaction("Alice", "Bob", 15)
	blockchain.Mining()
	blockchain.Print()
}
