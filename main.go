package main

import (
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("# You MUST supply 'path/to/file' and 'password' (which can be empty)")
		fmt.Println("# E.g. ethprivkey '/path/to/key.json' 'watwatchickabutt'")
		return
	}

	filePath := os.Args[1]
	password := os.Args[2]

	keyjson, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err)
	}

	key, err := keystore.DecryptKey(keyjson, password)
	if err != nil {
		panic(err)
	}
	privateKey := hex.EncodeToString(crypto.FromECDSA(key.PrivateKey))
	fmt.Println(privateKey)
}
