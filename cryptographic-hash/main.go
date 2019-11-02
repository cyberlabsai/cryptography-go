package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
)

// Hashs

func main() {

	// cryptographic hash
	s := "Cryptofriday 2 é o melhor evento hacker de Floripa!"

	md5 := md5.Sum([]byte(s))
	sha1 := sha1.Sum([]byte(s))
	sha256 := sha256.Sum256([]byte(s))

	fmt.Printf("MD5: %x\n", md5)
	fmt.Printf("SHA1: %x\n", sha1)
	fmt.Printf("SHA2-256: %x\n", sha256)

}
