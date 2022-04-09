package model

import (
	"crypto/sha256"
	"encoding/json"
	"log"
	"time"
)

type Block struct {
	Timestamp    int64
	Nonce        int
	PreviousHash [32]byte
	Transactions []*Transaction
}

func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	return &Block{
		Nonce:        nonce,
		PreviousHash: previousHash,
		Timestamp:    time.Now().UnixNano(),
		Transactions: transactions,
	}
}

func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Nonce        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previousHash"`
		Timestamp    int64          `json:"timestamp"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Nonce:        b.Nonce,
		PreviousHash: b.PreviousHash,
		Timestamp:    b.Timestamp,
		Transactions: b.Transactions,
	})
}

func (b *Block) Print() {
	log.Printf("Timestamp 		%d\n", b.Timestamp)
	log.Printf("Nonce 			%d\n", b.Nonce)
	log.Printf("Previous Hash 	%x\n", b.PreviousHash)
	log.Printf("Transactions 	%s\n", b.Transactions)
}

func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}
