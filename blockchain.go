package main

import (
	"blockchain/model"
	"log"
	"strings"
)

type Blockchain struct {
	transactionPool []*model.Transaction
	chain           []*model.Block
}

func NewBlockchain() *Blockchain {
	b := &model.Block{}
	bc := &Blockchain{}
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		log.Printf("%s Chain %d %s \n", strings.Repeat("-", 25), i, strings.Repeat("-", 25))
		block.Print()
	}
	log.Printf("%s\n", strings.Repeat("-", 25))
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *model.Block {
	b := model.NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*model.Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *model.Block {
	return bc.chain[len(bc.chain)-1]
}

func (bc Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := model.NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

func init() {
	log.SetPrefix("Blockchain: ")
}

func main() {
	blockchain := NewBlockchain()
	blockchain.Print()

	blockchain.AddTransaction("Alice", "Bob", 5)
	previousHash := blockchain.LastBlock().Hash()
	blockchain.CreateBlock(5, previousHash)
	blockchain.Print()

	blockchain.AddTransaction("Bob", "Alice", 10)
	blockchain.AddTransaction("Alice", "Bob", 15)
	previousHash = blockchain.LastBlock().Hash()
	blockchain.CreateBlock(10, previousHash)
	blockchain.Print()
}
