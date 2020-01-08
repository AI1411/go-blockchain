package main

import (
	"crypto/sha256"
	"encoding/json"
	"fmt"
	"log"
	"strings"
	"time"
)

//先頭に３つ0が続くように設定
const (
	MINING_DIFICULTY = 3
	MINING_SENDER    = "THE BLOCKCHAIN"
	MINING_REWARD    = 1.0
)

//ブロックのストラクトを定義
type Block struct {
	timestamp    int64
	nonce        int
	previousHash [32]byte
	transactions []*Transaction
}

//新しいブロックを作成
func NewBlock(nonce int, previousHash [32]byte, transactions []*Transaction) *Block {
	b := new(Block)
	b.timestamp = time.Now().UnixNano()
	b.nonce = nonce
	b.previousHash = previousHash
	b.transactions = transactions
	return b
}

//ブロックの中身を出力
func (b *Block) Print() {
	fmt.Printf("timestamp         %d\n", b.timestamp)
	fmt.Printf("nonce             %d\n", b.nonce)
	fmt.Printf("previousHash      %x\n", b.previousHash)
	for _, t := range b.transactions {
		t.Print()
	}
}

//トランザクションのストラクトを定義
type Transaction struct {
	senderBlockchainAddress    string
	recipientBlockchainAddress string
	value                      float32
}

//トランザクションを新規作成
func NewTransaction(sender string, recipient string, value float32) *Transaction {
	return &Transaction{sender, recipient, value}
}

//トランザクションを出力
func (t *Transaction) Print() {
	fmt.Printf("%s\n", strings.Repeat("-", 40))
	fmt.Printf(" sender_blockchain_address      %s\n", t.senderBlockchainAddress)
	fmt.Printf(" recipient_blockchain_address   %s\n", t.recipientBlockchainAddress)
	fmt.Printf(" value                          %.1f\n", t.value)
}

//jsonmarshalに変換
func (t *Transaction) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Sender    string  `json:"sender_blockchain_address"`
		Recipient string  `json:"recipient_blockchain_address"`
		Value     float32 `json:"value"`
	}{
		Sender:    t.senderBlockchainAddress,
		Recipient: t.recipientBlockchainAddress,
		Value:     t.value,
	})
}

//ブロックチェーンのストラクト作成
type Blockchain struct {
	transactionPool   []*Transaction
	chain             []*Block
	blockchainAddress string
}

//ブロックチェーンを新規作成
func NewBlockchain(blockchainAddress string) *Blockchain {
	b := &Block{}
	bc := new(Blockchain)
	bc.blockchainAddress = blockchainAddress
	bc.CreateBlock(0, b.Hash())
	return bc
}

//作成したブロックをプールに追加していく
func (bc *Blockchain) CreateBlock(nonce int, previousHash [32]byte) *Block {
	b := NewBlock(nonce, previousHash, bc.transactionPool)
	bc.chain = append(bc.chain, b)
	bc.transactionPool = []*Transaction{}
	return b
}

//最後のチェーン
func (bc *Blockchain) LastBlock() *Block {
	return bc.chain[len(bc.chain)-1]
}

//出力が見やすいように整形
func (bc *Blockchain) Print() {
	for i, block := range bc.chain {
		fmt.Printf("%s Chain %d %s\n", strings.Repeat("=", 25), i,
			strings.Repeat("=", 25))
		block.Print()
	}
	fmt.Printf("%s\n", strings.Repeat("*", 25))
}

//トランザクション追加のメソッド
func (bc *Blockchain) AddTransaction(sender string, recipient string, value float32) {
	t := NewTransaction(sender, recipient, value)
	bc.transactionPool = append(bc.transactionPool, t)
}

//先頭三文字がゼロになるかどうかを判定するメソッド
func (bc *Blockchain) ValidProof(nonce int, previousHash [32]byte, transactions []*Transaction, difficulty int) bool {
	zeros := strings.Repeat("0", difficulty)
	guessBlock := Block{0, nonce, previousHash, transactions}
	guessHashStr := fmt.Sprintf("%x", guessBlock.Hash())
	//fmt.Println(guessHashStr)
	return guessHashStr[:difficulty] == zeros
}

//マイニング
func (bc *Blockchain) Mining() bool {
	bc.AddTransaction(MINING_SENDER, bc.blockchainAddress, MINING_REWARD)
	//rewardが入った状態でproof of workをする
	nonce := bc.ProofOfWork()
	previousHash := bc.LastBlock().Hash()
	//nonceを含めて新しいブロックを生成する
	bc.CreateBlock(nonce, previousHash)
	log.Println("action=mining, status=success")
	return true
}

//valueの計算
func (bc *Blockchain) CalculateTotalAmount(blockchainAddress string) float32 {
	var totalAmount float32 = 0.0
	for _, b := range bc.chain {
		for _, t := range b.transactions {
			value := t.value
			//受取人だったらvalueが増える
			if blockchainAddress == t.recipientBlockchainAddress {
				totalAmount += value
			}
			//送り手だった場合valueが減る
			if blockchainAddress == t.senderBlockchainAddress {
				totalAmount -= value
			}
		}
	}
	return totalAmount
}

func (bc *Blockchain) ProofOfWork() int {
	transactions := bc.CopyTransactionPool()
	previousHash := bc.LastBlock().Hash()
	//nonceの初期化
	nonce := 0
	//ValidProofがtrueになるまでループで回す
	for !bc.ValidProof(nonce, previousHash, transactions, MINING_DIFICULTY) {
		nonce += 1
	}
	return nonce
}

//トランザクションプールをコピー
func (bc *Blockchain) CopyTransactionPool() []*Transaction {
	transactions := make([]*Transaction, 0)
	for _, t := range bc.transactionPool {
		transactions = append(transactions, NewTransaction(t.senderBlockchainAddress, t.recipientBlockchainAddress, t.value))
	}
	return transactions
}

//sha256でハッシュ化
func (b *Block) Hash() [32]byte {
	m, _ := json.Marshal(b)
	return sha256.Sum256([]byte(m))
}

//jsonを上書き
func (b *Block) MarshalJSON() ([]byte, error) {
	return json.Marshal(struct {
		Timestamp    int64          `json:"timestamp"` //struct tag
		Nance        int            `json:"nonce"`
		PreviousHash [32]byte       `json:"previous_hash"`
		Transactions []*Transaction `json:"transactions"`
	}{
		Timestamp:    b.timestamp,
		Nance:        b.nonce,
		PreviousHash: b.previousHash,
		Transactions: b.transactions,
	})
}

func init() {
	log.SetPrefix("Blockchain")
}
func main() {
	myBlockchainAddress := "my_blockchain_address"
	blockChain := NewBlockchain(myBlockchainAddress)
	blockChain.Print()

	//AさんがBさんに1.0のvalueを送る
	blockChain.AddTransaction("A", "B", 1.0)
	//マイニング
	blockChain.Mining()
	blockChain.Print()

	blockChain.AddTransaction("C", "D", 3.0)
	blockChain.Mining()
	blockChain.Print()

	fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
	fmt.Printf("A %.1f\n", blockChain.CalculateTotalAmount("B"))
	fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("D"))
}
