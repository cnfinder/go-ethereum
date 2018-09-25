package main

import (
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/base58"
)

func main() {
	//Create an account

	fmt.Printf("bumeng: %d %v\n", len("bubiV8i7c3fM7MPFfjZsP4h1zGLUEix4gXUnX8Ft"), "bubiV8i7c3fM7MPFfjZsP4h1zGLUEix4gXUnX8Ft")
	key, err := crypto.GenerateKey()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	}

	//Get the address
	address := crypto.PubkeyToAddress(key.PublicKey).Hex()
	fmt.Printf("address[%d][%v]\n", len(address), address)

	//Get the private key
	privateKey := hex.EncodeToString(key.D.Bytes())
	fmt.Printf("privateKey[%d][%v]\n", len(privateKey), privateKey)
}
