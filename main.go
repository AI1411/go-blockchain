package main

import (
	"blockchain/model"
	"log"
)

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := model.NewBlockchain()
	blockchain.Print()

	blockchain.AddTransaction("Alice", "Bob", 5)
	previousHash := blockchain.LastBlock().Hash()
	nonce := blockchain.ProofOfWork()
	blockchain.CreateBlock(nonce, previousHash)
	blockchain.Print()

	blockchain.AddTransaction("Bob", "Alice", 10)
	blockchain.AddTransaction("Alice", "Bob", 15)
	previousHash = blockchain.LastBlock().Hash()
	nonce = blockchain.ProofOfWork()
	blockchain.CreateBlock(nonce, previousHash)
	blockchain.Print()
}
