package main

import (
	"fmt"
	secp "github.com/cosmos/cosmos-sdk/crypto/keys/secp256k1"
	eSecp "github.com/ethereum/go-ethereum/crypto/secp256k1"
)

func main() {
	sk := secp.GenPrivKey()
	fmt.Println(sk.String())
	msg := []byte("Hello")
	sig, err := sk.Sign(msg)
	if err != nil {
		panic(err)
	}
	fmt.Println(sig)
	ok := eSecp.VerifySignature(sk.PubKey().Bytes(), msg, sig)
	if !ok {
		fmt.Println("could not verify")
		return
	}
	fmt.Println("verification ok")
}
