package wallet

import (
	"encoding/base64"
	"fmt"

	// Use a cryptography library.

	"github.com/99designs/keyring"
	"github.com/cossacklabs/themis/gothemis/keys"
	"github.com/kevinburke/nacl"
	"github.com/kevinburke/nacl/secretbox"
)

// themis
func createKeyEC() *keys.Keypair {
	keyPair, err := keys.New(keys.KEYTYPE_EC)
	if nil != err {
		fmt.Println("Keypair generating error")
		return
	}
	fmt.Println(base64.StdEncoding.EncodeToString(keyPair.Private.Value))
	fmt.Println(base64.StdEncoding.EncodeToString(keyPair.Public.Value))
	return (keyPair)
}

// Nacl
func callNacl() {
	key, err := nacl.Load("6368616e676520746869732070617373776f726420746f206120736563726574")
	if err != nil {
		panic(err)
	}
	encrypted := secretbox.EasySeal([]byte("hello world"), key)
	fmt.Println(base64.StdEncoding.EncodeToString(encrypted))
}

// Keyring
func Hello() string {

	ring, _ := keyring.Open(keyring.Config{
		ServiceName: "example",
	})

	_ = ring.Set(keyring.Item{
		Key:  "foo",
		Data: []byte("secret-bar"),
	})

	i, _ := ring.Get("foo")

	fmt.Printf("%s", i.Data)

	return ""
}
