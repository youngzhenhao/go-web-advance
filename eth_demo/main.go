package main

import (
	"flag"
	"fmt"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	var pkStr string
	flag.StringVar(&pkStr, "p", "", "private key")
	flag.Parse()
	if flag.NFlag() == 0 {
		flag.Usage()
		return
	}
	address, err := PrivateKeyToAddress(pkStr)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(address)
}

func PrivateKeyToAddress(pkStr string) (string, error) {
	privateKey, err := crypto.HexToECDSA(pkStr)
	if err != nil {
		return "", err
	}
	address := crypto.PubkeyToAddress(privateKey.PublicKey).Hex()
	return address, nil
}
