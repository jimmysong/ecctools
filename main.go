package main

import (
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/btcsuite/btcd/btcec"
)

func fromHex(s string) *big.Int {
	r, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("invalid hex in source file: " + s)
	}
	return r
}

func toHex(i *big.Int) string {
	return hex.EncodeToString(i.Bytes())
}

func main() {
	secp256k1 := btcec.S256()
	operation := os.Args[1]
	if operation == "add" {
		x1, y1 := fromHex(os.Args[2]), fromHex(os.Args[3])
		x2, y2 := fromHex(os.Args[4]), fromHex(os.Args[5])
		rx, ry := secp256k1.Add(x1, y1, x2, y2)
		fmt.Printf("%v %v\n", toHex(rx), toHex(ry))
	} else if operation == "mult" {
		k, x1, y1 := fromHex(os.Args[2]).Bytes(), fromHex(os.Args[3]), fromHex(os.Args[4])
		rx, ry := secp256k1.ScalarMult(x1, y1, k)
		fmt.Printf("%v %v\n", toHex(rx), toHex(ry))
	} else if operation == "multg" {
		k := fromHex(os.Args[2]).Bytes()
		rx, ry := secp256k1.ScalarBaseMult(k)
		fmt.Printf("%v %v\n", toHex(rx), toHex(ry))
	}
}
