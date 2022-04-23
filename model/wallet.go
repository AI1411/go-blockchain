package model

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"github.com/btcsuite/btcutil/base58"
	"golang.org/x/crypto/ripemd160"
)

type Wallet struct {
	PrivateKey        *ecdsa.PrivateKey
	PublicKey         *ecdsa.PublicKey
	BlockchainAddress string
}

func NewWallet() *Wallet {
	w := new(Wallet)
	// 公開鍵(64byte)、秘密鍵(32byte)生成
	privateKey, _ := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	w.PrivateKey = privateKey
	w.PublicKey = &w.PrivateKey.PublicKey

	h2 := sha256.New()
	h2.Write(w.PublicKey.X.Bytes())
	h2.Write(w.PublicKey.Y.Bytes())
	digest2 := h2.Sum(nil)

	h3 := ripemd160.New()
	h3.Write(digest2)
	digest3 := h3.Sum(nil)

	vd4 := make([]byte, 21)
	vd4[0] = 0x00
	copy(vd4[1:], digest3[:])

	h5 := sha256.New()
	h5.Write(vd4)
	digest5 := h5.Sum(nil)

	h6 := sha256.New()
	h6.Write(digest5)
	digest6 := h6.Sum(nil)

	cSum := digest6[:4]
	dc8 := make([]byte, 25)
	copy(dc8[:21], vd4[:])
	copy(dc8[21:], cSum[:])

	address := base58.Encode(dc8)
	w.BlockchainAddress = address

	return w
}

func (w *Wallet) PrintPrivateKeyStr() string {
	return fmt.Sprintf("%x", w.PrivateKey.D.Bytes())
}

func (w *Wallet) PrintPublicKeyStr() string {
	return fmt.Sprintf("%x%x", w.PublicKey.X.Bytes(), w.PublicKey.Y.Bytes())
}
