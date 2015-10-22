package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"math/big"
	"os"

	"github.com/btcsuite/btcd/btcec"
)

func fromHex(s string) *big.Int {
	r, ok := new(big.Int).SetString(s, 16)
	if !ok {
		panic("invalid hex: " + s)
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
	} else if operation == "sadd" {
		k1, k2 := fromHex(os.Args[2]), fromHex(os.Args[3])
		rk := new(big.Int)
		rk.Add(k1, k2)
		rk.Mod(rk, secp256k1.N)
		fmt.Printf("%x\n", rk.Bytes())
	} else if operation == "smult" {
		k1, k2 := fromHex(os.Args[2]), fromHex(os.Args[3])
		rk := new(big.Int)
		rk.Mul(k1, k2)
		rk.Mod(rk, secp256k1.N)
		fmt.Printf("%x\n", rk.Bytes())
	} else if operation == "mult" {
		k, x1, y1 := fromHex(os.Args[2]).Bytes(), fromHex(os.Args[3]), fromHex(os.Args[4])
		rx, ry := secp256k1.ScalarMult(x1, y1, k)
		fmt.Printf("%v %v\n", toHex(rx), toHex(ry))
	} else if operation == "multg" {
		k := fromHex(os.Args[2]).Bytes()
		rx, ry := secp256k1.ScalarBaseMult(k)
		fmt.Printf("%v %v\n", toHex(rx), toHex(ry))
	} else if operation == "sha256" {
		s := fromHex(os.Args[2]).Bytes()
		fmt.Printf("%x\n", sha256.Sum256(s))
	}
}
