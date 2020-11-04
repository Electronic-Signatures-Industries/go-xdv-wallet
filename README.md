# go-xdv-wallet
Next Generation Cryptographic Wallet for Documents, Blockchains and Verifiable Credentials

## Example
```golang
package main

import (
	"bytes"
	"crypto"
	"golang.org/x/crypto/ed25519"
	"log"
)

type zeroReader struct{}

func (zeroReader) Read(buf []byte) (int, error) {
	for i := range buf {
		buf[i] = 0
	}
	return len(buf), nil
}

func TestSignVerify() {
	var zero zeroReader
	public, private, _ := ed25519.GenerateKey(zero)
	
	message := []byte("test message")
	sig := ed25519.Sign(private, message)
	log.Println(ed25519.Verify(public, message, sig))
	
	wrongMessage := []byte("wrong message")
	log.Println(ed25519.Verify(public, wrongMessage, sig))
}

func TestCryptoSigner() {
	var zero zeroReader
	public, private, _ := ed25519.GenerateKey(zero)
	
	signer := crypto.Signer(private)
	
	publicInterface := signer.Public()
	public2, ok := publicInterface.(ed25519.PublicKey)
	if !ok {
		log.Fatalf("expected PublicKey from Public() but got %T", publicInterface)
	}
	
	if !bytes.Equal(public, public2) {
		log.Fatalf("public keys do not match: original:%x vs Public():%x", public, public2)
	}
	
	message := []byte("message")
	var noHash crypto.Hash
	signature, err := signer.Sign(zero, message, noHash)
	if err != nil {
		log.Fatalf("error from Sign(): %s", err)
	}
	
	log.Println(ed25519.Verify(public, message, signature))
}

func main() {
	TestSignVerify()
	TestCryptoSigner()
}
```
