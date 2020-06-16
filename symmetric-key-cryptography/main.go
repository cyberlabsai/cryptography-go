package main

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
	"fmt"
)

// AES-GCM should be used because the operation is an authenticated encryption
// algorithm designed to provide both data authenticity (integrity) as well as
// confidentiality.

func encrypt() {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("HackweekTheBest!")
	plaintext := []byte("Este é o texto plano a ser cifrado")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := []byte("Hackweek!!!!")

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Printf("Texto cifrado: %x\n", ciphertext)
}

func decrypt() {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("HackweekTheBest!")
	ciphertext, _ := hex.DecodeString("b5b80251d90d055f8cc2de7cbdeeb4237f5956981c8132fe6602541450cbbdfad7626c2767b41ddfe3852d909fe3d338484708")

	nonce := []byte("Hackweek!!!!")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	plaintext, err := aesgcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Printf("Texto plano original: %s\n", string(plaintext))
}

func main() {

	// cifrar texto

	encrypt()

	// decifrar texto

	decrypt()
}
