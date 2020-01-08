package wallet

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"fmt"
)

//ウォレットのストラクトを定義
type Wallet struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
}

//ウォレットを新規作成
func NewWallet() *Wallet {
	w := new(Wallet)
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.privateKey = privateKey
	w.publicKey = &w.privateKey.PublicKey
	return w
}

//privateKeyを生成
func (w *Wallet) PrivateKey() *ecdsa.PrivateKey {
	return w.privateKey
}

//privateKeyの文字列を取得
func (w *Wallet) PrivateKeyStr() string  {
	return fmt.Sprintf("%x", w.privateKey.D.Bytes())
}

//privateKeyを生成
func (w *Wallet) PublicKey() *ecdsa.PublicKey {
	return w.publicKey
}

//publicKeyの文字列を取得
func (w *Wallet) PublicKeyStr() string  {
	return fmt.Sprintf("%x%x", w.publicKey.X.Bytes(), w.publicKey.Y.Bytes())
}