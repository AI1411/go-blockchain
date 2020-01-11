package main

import (
	"html/template"
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

//wallet server起動
func (ws *WalletServer) Run() {
	http.HandleFunc("/", ws.Index)
	log.Fatal(http.ListenAndServe("0.0.0.0:"+strconv.Itoa(int(ws.Port())), nil))
}
