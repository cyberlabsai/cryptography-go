package main

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

// RSA

const (
	rsaKeySize = 2048
)

type keypair struct {
	priv *rsa.PrivateKey
	pub  *rsa.PublicKey
}

var kp keypair
var ciphertext, signedMessage []byte
var rng io.Reader

func generateKeypair() error {
	var err error
	kp.priv, err = rsa.GenerateKey(rand.Reader, rsaKeySize)
	if err != nil {
		return err
	}
	kp.pub = &kp.priv.PublicKey
	return nil
}

func encrypt() {

	var err error

	secretMessage := []byte("criptofryday2 is the best hacker event in Florianópolis")
	label := []byte("talk")

	ciphertext, err = rsa.EncryptOAEP(sha256.New(), rng, kp.pub, secretMessage, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return
	}

	// Since encryption is a randomized function, ciphertext will be
	// different each time.
	fmt.Printf("Ciphertext: %x\n", ciphertext)
}

func decrypt() {

	label := []byte("talk")

	plaintext, err := rsa.DecryptOAEP(sha256.New(), rng, kp.priv, ciphertext, label)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return
	}

	fmt.Printf("Plaintext: %s\n", string(plaintext))
}

func sign() {
	var err error
	message := []byte("criptofryday2 is the best hacker event in Florianópolis")
	signedMessage, err = rsa.EncryptPKCS1v15(rng, kp.pub, message)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from encryption: %s\n", err)
		return
	}
	fmt.Printf("Mensagem assinada: %x\n", signedMessage)
}

func verify() {

	msgVerified, err := rsa.DecryptPKCS1v15(rng, kp.priv, signedMessage)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error from decryption: %s\n", err)
		return
	}

	fmt.Printf("Mensagem Verificada: %s\n", string(msgVerified))
}

func main() {

	// crypto/rand.Reader is a good source of entropy for randomizing the
	// encryption function.
	rng = rand.Reader

	// gerar par de chaves
	generateKeypair()

	// cifrar texto
	encrypt()

	// decifrar texto
	decrypt()

	// assinar mensagem
	sign()

	// verificar mensagem
	verify()
}
