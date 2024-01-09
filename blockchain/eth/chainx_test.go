package eth

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGenerate(t *testing.T) {
	var acc Account
	acc = Account{}
	addr, priKey, pubKey, _ := acc.Generate()
	fmt.Println(priKey)
	fmt.Println(pubKey)
	fmt.Println(addr)

}

func TestGeneratePubKey(t *testing.T) {
	var acc Account = Account{}
	priKey := "0x27180f313983df4d4e5bb97e6f951a8d94f03af96ba82ade224df156b290f3e0"
	priKey2 := "27180f313983df4d4e5bb97e6f951a8d94f03af96ba82ade224df156b290f3e0"
	pubKey := "0x04f75ecc85fe618ac381375bcd503942a4b0f462e2215172d380768cbbaef17a8faae77099a1f324c9154ab691577c135899a3ec698d651b3bab8077fb57ac2005"
	addr := "0x7C0d88BAf40bfC15a6672f5bB79AA1825cFAd005"
	address, publicKey, err := acc.GeneratePubKey(priKey)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, pubKey, publicKey)
	assert.Equal(t, addr, address)

	address, publicKey, err = acc.GeneratePubKey(priKey2)
	if err != nil {
		t.Error(err)
	}
	assert.Equal(t, pubKey, publicKey)
	assert.Equal(t, addr, address)
}

func TestEvmPrivateKeyToIAA(t *testing.T) {
	var acc = Account{}
	priKeyEVM := "27180f313983df4d4e5bb97e6f951a8d94f03af96ba82ade224df156b290f3e0"
	privateKeyIAA := "0xfcd2efcc2027180f313983df4d4e5bb97e6f951a8d94f03af96ba82ade224df156b290f3e0"
	pubKeyIAA := "0xf3b3cd032103f75ecc85fe618ac381375bcd503942a4b0f462e2215172d380768cbbaef17a8f"
	priKeyIAA, pubKeyIAAByte, err := acc.EvmPrivateKeyToIAA(priKeyEVM)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("privateKeyIAA: \t", priKeyIAA)
	fmt.Println("pubKeyIAA: \t\t", pubKeyIAAByte)

	assert.Equal(t, privateKeyIAA, privateKeyIAA)
	assert.Equal(t, pubKeyIAA, pubKeyIAAByte)
}

func TestHDWalletGenerate(t *testing.T) {
	var acc = Account{}
	nmemonic, privateKeyHex, publicKeyHex, address, err := acc.HDWalletGenerate()
	if err != nil {
		t.Error(err)
	}
	fmt.Println("nmemonic:\t", nmemonic)
	fmt.Println("privateKeyHex:\t", privateKeyHex)
	fmt.Println("publicKeyHex:\t", publicKeyHex)
	fmt.Println("address:\t", address)

}

func TestHDWalletGenerateWithNmemonic(t *testing.T) {
	var acc = Account{}
	nmemonic := "scissors patrol tongue twin february blouse betray coyote mixture nice inside antenna"
	privateKeyExp := "04366de1c7e19eb91efedac684cb6951560052d701ff9068cb292d6f4ca98c1e"
	publicKeyExp := "48ed9af0b8b691fe3c447d076f15ccfa1ab5205c480d193e0ba92c2787ec479a6ef812d66c5ba9e61e2652e4b7f194285decd0ed169a4875f18df4aa64736ee9"
	addressExp := "0xE863D7551AF81c070B9521ABECD76e8BCdAF4d04"
	privateKeyHex, publicKeyHex, address, err := acc.HDWalletGenerateWithNmemonic(nmemonic)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("privateKeyHex:\t", privateKeyHex)
	fmt.Println("publicKeyHex:\t", publicKeyHex)
	fmt.Println("address:\t", address)
	assert.Equal(t, privateKeyExp, privateKeyHex)
	assert.Equal(t, publicKeyExp, publicKeyHex)
	assert.Equal(t, addressExp, address)

	nmemonic24 := "street glance high ski dinner mixture topple meadow mechanic buddy during bridge vendor buyer mass remember identify jeans about grocery talk piano palace negative"
	//nmemonic24 := "sorry liberty pioneer weird senior jungle author canvas girl patrol wrestle magnet struggle shop turn trial rate wrist trouble uncover lend enough cushion bless"

	privateKeyExp = "61decf8be7824f624a1513d3c4d0ce3af60851c2501a201e239fa19c90efdfcd"
	publicKeyExp = "97773d139ef556af64c36937abf952f02e0544f52e32b8108092e155d4f4f79df178a9afa583f49ee1db25162072a841484c21938094176a80b7a56738180621"
	addressExp = "0xDfb44BDe559fFe1ed37ba3F46a82F1353795c4bF"
	privateKeyHex, publicKeyHex, address, err = acc.HDWalletGenerateWithNmemonic(nmemonic24)
	if err != nil {
		t.Error(err)
	}
	fmt.Println("privateKeyHex24:\t", privateKeyHex)
	fmt.Println("publicKeyHex24:\t", publicKeyHex)
	fmt.Println("address24:\t", address)
	assert.Equal(t, privateKeyExp, privateKeyHex)
	assert.Equal(t, publicKeyExp, publicKeyHex)
	assert.Equal(t, addressExp, address)
}

func TestA(t *testing.T) {

}
