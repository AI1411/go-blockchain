package main

import (
	"blockchain/wallet"
	"fmt"
	"log"
)

func init() {
log.SetPrefix("Blockchain")
}
func main() {
	w := wallet.NewWallet()
	fmt.Println(w.PrivateKey())
	fmt.Println(w.PublicKey())

	fmt.Println(w.BlockchainAddress())
//myBlockchainAddress := "my_blockchain_address"
//blockChain := NewBlockchain(myBlockchainAddress)
//blockChain.Print()
//
////AさんがBさんに1.0のvalueを送る
//blockChain.AddTransaction("A", "B", 1.0)
////マイニング
//blockChain.Mining()
//blockChain.Print()
//
//blockChain.AddTransaction("C", "D", 3.0)
//blockChain.Mining()
//blockChain.Print()
//
//fmt.Printf("my %.1f\n", blockChain.CalculateTotalAmount("my_blockchain_address"))
//fmt.Printf("A %.1f\n", blockChain.CalculateTotalAmount("B"))
//fmt.Printf("C %.1f\n", blockChain.CalculateTotalAmount("D"))
}
