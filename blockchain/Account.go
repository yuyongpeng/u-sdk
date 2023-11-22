package blockchain

import "github.com/tyler-smith/go-bip39"

/**
用于处理通用的eth和irita的函数
*/

/*
*
生成12个助记词 以太坊用
*/
func GenerateMnemonic12() (mnemonic string, err error) {
	return generateMnemonic(128)
}

/*
*
生成24个助记词  irita使用
*/
func GenerateMnemonic24() (nmemonic string, err error) {
	return generateMnemonic(256)
}

/*
*
bitSize = 128 / 256
*/
func generateMnemonic(bitSize int) (monemonic string, err error) {
	rand, err := bip39.NewEntropy(bitSize)
	if err != nil {
		return "", err
	}
	mnemonic, err := bip39.NewMnemonic(rand)
	if err != nil {
		return "", err
	}
	return mnemonic, nil
}
