package irita

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestBech32ToHex(t *testing.T) {
	addressIAA := "iaa1phqftjq49s35tm73g8key5j9ut8jrhuux2hr6h"
	addressEVM := "0x0DC095C8152C2345EFD141ED925245E2CF21DF9C"
	address, err := Bech32ToHex(addressIAA, "iaa")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("address: ", addressIAA)
	assert.Equal(t, addressEVM, address)
}

func TestHexToBech32(t *testing.T) {
	addressIAA := "iaa1phqftjq49s35tm73g8key5j9ut8jrhuux2hr6h"
	addressEVM := "0x0DC095C8152C2345EFD141ED925245E2CF21DF9C"
	address, err := HexToBech32(addressEVM)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("iaa: ", address)
	assert.Equal(t, addressIAA, address)
}

func TestGetFromBech32(t *testing.T) {

	//str := "world"
	////x := hex.EncodeToString(([]byte)(str))
	//fmt.Printf("%X", str) // 616263646566  hex
	bStr := "abcdef"
	bBytes := []byte(bStr)
	fmt.Printf("%x", bBytes)
}
