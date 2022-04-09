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
	w := model.NewWallet()
	fmt.Println(w.PrintPrivateKeyStr())
	fmt.Println(w.PrintPublicKeyStr())
}
