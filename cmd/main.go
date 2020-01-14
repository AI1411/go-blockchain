package main

import (
	"blockchain/utils"
	"fmt"
)

func main() {
	//接続確認 bool
	fmt.Println(utils.FindNeighbors("127.0.0.1", 5000, 0,3,5000,5003))
}
