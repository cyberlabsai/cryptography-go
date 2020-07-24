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
	key := []byte("CybertalksTalk!!")
	plaintext := []byte("Este é o texto plano a ser cifrado")

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err.Error())
	}

	// Never use more than 2^32 random nonces with a given key because of the risk of a repeat.
	nonce := []byte("Cybertalks!!")

	aesgcm, err := cipher.NewGCM(block)
	if err != nil {
		panic(err.Error())
	}

	ciphertext := aesgcm.Seal(nil, nonce, plaintext, nil)
	fmt.Println(" ")
	fmt.Printf("Texto cifrado: %x\n", ciphertext)
}

func decrypt() {
	// The key argument should be the AES key, either 16 or 32 bytes
	// to select AES-128 or AES-256.
	key := []byte("CybertalksTalk!!")
	ciphertext, _ := hex.DecodeString("4ba79451b86fc30e4615e0d2d5524121938fb9577e8fef49ba219acf233ecc018f4aed81ca3cde6b7ba708b56af7958d9970d7")

	nonce := []byte("Cybertalks!!")

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

	fmt.Println(" ")
	fmt.Printf("Texto plano original: %s\n", string(plaintext))

}

func main() {

	fmt.Println("Criptografia SIMÉTRICA")

	// cifrar texto

	encrypt()

	// decifrar texto

	decrypt()
}
