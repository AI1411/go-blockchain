package model

import (
	"fmt"
	"log"
	"strings"
)

const MiningDifficulty = 3

type Blockchain struct {
	TransactionPool []*Transaction
	Chain           []*Block
}

func NewBlockchain() *Blockchain {
	b := &Block{}
	bc := &Blockchain{}
	bc.CreateBlock(0, b.Hash())
	return bc
}

func (bc *Blockchain) Print() {
	for i, block := range bc.Chain {
		log.Printf("%s Chain %d %s \n", strings.Repeat("-", 25), i, strings.Repeat("-", 25))
		block.Print()
	}
	log.Printf("%s\n", strings.Repeat("-", 25))
}

func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.TransactionPool)
	bc.Chain = append(bc.Chain, b)
	bc.TransactionPool = []*Transaction{}
	return b
}

func (bc *Blockchain) LastBlock() *Block {
	return bc.Chain[len(bc.Chain)-1]
}

func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.TransactionPool = append(bc.TransactionPool, t)
}

func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, len(bc.TransactionPool))
	for _, t := range bc.TransactionPool {
		transactions = append(transactions,
			NewTransaction(t.SenderBlockchainAddress,
				t.RecipientBlockchainAddress,
				t.Value))
	}
	return transactions
}

func (bc Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	return guessHashStr[:difficulty] == zeros
}

func (bc Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	nonce := 0
	for !bc.ValidProof(nonce, previousHash, transactions, MiningDifficulty) {
		nonce += 1
	}
	return nonce
}
