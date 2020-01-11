package main

import (
	"blockchain/utils"
	"blockchain/wallet"
	"html/template"
	"io"
	"log"
	"net/http"
	"path"
	"strconv"
)

const tempDir = "wallet_server/templates"

type WalletServer struct {
	port    uint16
	gateway string
}

//wallet serverを新規作成
func NewWalletServer(port uint16, gateway string) *WalletServer {
	return &WalletServer{port, gateway}
}

//portを取得
func (ws *WalletServer) Port() uint16 {
	return ws.port
}

//Gatewayを取得
func (ws *WalletServer) Gateway() string {
	return ws.gateway
}

//templateの読み込み
func (ws *WalletServer) Index(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		t, _ := template.ParseFiles(path.Join(tempDir, "index.html"))
		t.Execute(w, "")
	default:
		log.Printf("Error:Invalid HTTP Method")

	}
}

func (ws *WalletServer) Wallet(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		w.Header().Add("Content-type", "application/json")
		myWallet := wallet.NewWallet()
		m,_ := myWallet.MarshalJSON()
		io.WriteString(w, string((m[:])))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error: invalid http method")
	}
}

func (ws *WalletServer) CreateTransaction(w http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodPost:
		io.WriteString(w, string(utils.JsonStatus("success")))
	default:
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Error: invalid http methods")
	}
}

//wallet server起動
func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	http.HandleFunc("/wallet", ws.Wallet)
	http.HandleFunc("/transaction", ws.CreateTransaction)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
