package main

import (
	"blockchain/block"
	"blockchain/wallet"
	"io"
	"log"
	"net/http"
	"strconv"
)

var cache map[string]*block.Blockchain = make(map[string]*block.Blockchain)

//BlockchainServerのストラクトを定義
type BlockchainServer struct {
	port uint16
}

//ブロックチェーンサーバーを新規作成
func NewBlockchainServer(port uint16) *BlockchainServer {
	return &BlockchainServer{port}
}

//Portを指定
func (bcs *BlockchainServer) Port() uint16 {
	return bcs.port
}

//BlockChainを取得するメソッド
func (bcs *BlockchainServer) GetBlockchain() *block.Blockchain {
	bc, ok := cache["blockchain"]
	if !ok {
		minersWallet := wallet.NewWallet()
		bc = block.NewBlockchain(minersWallet.BlockchainAddress(), bcs.Port())
		//生成したブロックチェーンをキャッシュにいれる
		cache["blockchain"] = bc
		//確認用
		log.Printf("private_key %v", minersWallet.PrivateKeyStr())
		log.Printf("public_key %v", minersWallet.PublicKeyStr())
		log.Printf("private_key %v", minersWallet.BlockchainAddress())
	}
	return bc
}

func (bcs *BlockchainServer) GetChain(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		w.Header().Add("Content-Type", "application/json")
		bc := bcs.GetBlockchain()
		m, _ := bc.MarshalJSON()
		io.WriteString(w, string(m[:]))
	default:
		log.Printf("Error: Invalid http method")

	}
}

func (bcs *BlockchainServer) Run() {
	http.HandleFunc("/", bcs.GetChain)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(bcs.Port())), nil))
}
