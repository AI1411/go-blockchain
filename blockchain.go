package main

import (
	"blockchain/model"
	"crypto/sha256"
	"encoding/json"
	"log"
	"strings"
	"time"
)

type Block struct {
	nonce        int
	previousHash [32]byte
	timestamp    int64
	transactions []*model.Transaction
}

type Blockchain struct {
	transactionPool []*model.Transaction
	chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
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

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*model.Transaction{}
	return b
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*model.Transaction) *Block {
	return &Block{
		nonce:        nonce,
		previousHash: previousHash,
		timestamp:    time.Now().UnixNano(),
		transactions: transactions,
	}
}

func (b *Block) Print() {
	log.Printf("Timestamp 		%d\n", b.timestamp)
	log.Printf("Nonce 			%d\n", b.nonce)
	log.Printf("Previous Hash 	%x\n", b.previousHash)
	log.Printf("Transactions 	%s\n", b.transactions)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int                  `json:"nonce"`
		PreviousHash [32]byte             `json:"previousHash"`
		Timestamp    int64                `json:"timestamp"`
		Transactions []*model.Transaction `json:"transactions"`
	}{
		Nonce:        b.nonce,
		PreviousHash: b.previousHash,
		Timestamp:    b.timestamp,
		Transactions: b.transactions,
	})
}

func (bc *Blockchain) LastBlock() *Block {
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
