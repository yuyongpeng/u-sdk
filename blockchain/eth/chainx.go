package eth

import (
	"crypto/ecdsa"
	"errors"
	"fmt"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	iriscrypto "github.com/irisnet/core-sdk-go/common/crypto"
	cryptoamino "github.com/irisnet/core-sdk-go/common/crypto/codec"
	ethsecp256k1 "github.com/irisnet/core-sdk-go/common/crypto/keys/eth_secp256k1"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
	"log"
	"strings"
	"u-sdk/blockchain"
	"u-sdk/blockchain/irita"
)

type Account struct {
}

/*
生成以太坊的公私钥对
*/
func (acc Account) Generate() (address, priKey, pubKey string, err error) {
	privateKeyECDSA, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
		return "", "", "", err
	}
	// 私钥
	privateKeyBytes := crypto.FromECDSA(privateKeyECDSA)
	priKey = hexutil.Encode(privateKeyBytes)[2:]
	// 公钥
	pubKey2 := privateKeyECDSA.Public()
	publicKeyECDSA, ok := pubKey2.(*ecdsa.PublicKey)
	if !ok {
		errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return "", "", "", err
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey = hexutil.Encode(publicKeyBytes)
	// 地址
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()

	return
}

/*
*
根据EVM私钥生成公钥和address
*/
func (acc Account) GeneratePubKey(privKey string) (address, pubKey string, err error) {
	// 取消 hex 前缀
	privKey = strings.TrimLeft(privKey, "0x")
	privKey = strings.TrimLeft(privKey, "0X")

	privatekeyECDSA, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return "", "", err
	}
	// 公钥
	publicKey := privatekeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		errors.New("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
		return "", "", err
	}
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	pubKey = hexutil.Encode(publicKeyBytes)
	// 地址
	address = crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return
}

/*
EVM地址转换为原生的地址
*/
func (acc Account) AddressToIaa(address string) (addressIAA string, err error) {
	//Hex地址转换原生账户地址
	addressIAA, err = irita.HexToBech32(address)
	if err != nil {
		return "", err
	}
	return
}

// GetSecretKeyToCache
// 根据入参的evm 格式的私钥进行转换, 获取原生格式iaa 的私钥和公钥
func (acc Account) EvmPrivateKeyToIAA(priKey string) (privateKey, publicKey string, err error) {

	PrivateKey := &ethsecp256k1.PrivKey{
		Key: ethcommon.FromHex(priKey),
	}
	algo := "eth_secp256k1"
	armor := iriscrypto.EncryptArmorPrivKey(PrivateKey, "adfas", algo)
	km := iriscrypto.NewKeyManager()
	private, str, err := km.ImportPrivKey(armor, "adfas")
	fmt.Println(str)
	if err != nil {
		fmt.Println(err)
		return "", "", errors.New("import private key error")
	}
	privateKey = hexutil.Encode(cryptoamino.MarshalPrivKey(private))
	pubKey := km.ExportPubKey()
	if err != nil {
		return "", "", errors.New("encode private key error")
	}
	publicKey = hexutil.Encode(cryptoamino.MarshalPubkey(pubKey))

	return privateKey, publicKey, nil
}

/*
使用HDWallet 生成 公私钥信息
*/
func (acc Account) HDWalletGenerate() (nmemonic, privateKeyHex, publicKeyHex, address string, err error) {
	nmemonic, err = blockchain.GenerateMnemonic24()
	if err != nil {
		return "", "", "", "", err
	}

	privateKeyHex, publicKeyHex, address, err = acc.HDWalletGenerateWithNmemonic(nmemonic)
	if err != nil {
		return "", "", "", "", err
	}
	return
}

/*
使用 助记词 生成对应的私钥，公钥，地址
*/
func (acc Account) HDWalletGenerateWithNmemonic(nmemonic string) (privateKeyHex, publicKeyHex, address string, err error) {
	if err != nil {
		return "", "", "", err
	}
	path := "m/44'/60'/1'/0/0"
	wallet, err := hdwallet.NewFromMnemonic(nmemonic)
	if err != nil {
		return "", "", "", err
	}
	hdWalletPath := hdwallet.MustParseDerivationPath(path)
	account, err := wallet.Derive(hdWalletPath, false)
	if err != nil {
		return "", "", "", err
	}
	// 私钥 Hex
	privateKeyHex, err = wallet.PrivateKeyHex(account)
	if err != nil {
		return "", "", "", err
	}
	// 公钥 Hex
	publicKeyHex, err = wallet.PublicKeyHex(account)
	if err != nil {
		return "", "", "", err
	}
	// 地址
	address = account.Address.Hex()

	return
}
