package irita

import (
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"

	"github.com/btcsuite/btcutil/bech32"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// Bech32ToHex
// prefix 传入 iaa,将 iaa 开头的地址转换为 0x 格式的地址
func Bech32ToHex(addr, prefix string) (string, error) {
	bz, err := GetFromBech32(addr, prefix)
	if err != nil {
		return "", err
	}
	return common.BytesToAddress(bz).Hex(), nil
}

// evm 转 iaa 地址
func HexToBech32(addr string) (string, error) {
	bz, err := hexutil.Decode(addr)
	if err != nil {
		return "", err
	}
	t, err := ConvertAndEncode("iaa", bz)
	if err != nil {
		return "", err
	}
	return t, nil
}

// GetFromBech32 decodes a byte string from a Bech32 encoded string.
func GetFromBech32(bech32str, prefix string) ([]byte, error) {
	if len(bech32str) == 0 {
		return nil, errors.New("decoding Bech32 address failed: must provide an address")
	}

	hrp, bz, err := DecodeAndConvert(bech32str)
	if err != nil {
		return nil, err
	}

	if hrp != prefix {
		return nil, fmt.Errorf("invalid Bech32 prefix; expected %s, got %s", prefix, hrp)
	}

	return bz, nil
}

// ConvertAndEncode converts from a base64 encoded byte string to base32 encoded byte string and then to bech32
func ConvertAndEncode(hrp string, data []byte) (string, error) {
	converted, err := bech32.ConvertBits(data, 8, 5, true)
	if err != nil {
		return "", err
	}
	return bech32.Encode(hrp, converted)
}

// DecodeAndConvert decodes a bech32 encoded string and converts to base64 encoded bytes
func DecodeAndConvert(bech string) (string, []byte, error) {
	hrp, data, err := bech32.Decode(bech)
	if err != nil {
		return "", nil, err
	}
	converted, err := bech32.ConvertBits(data, 5, 8, false)
	if err != nil {
		return "", nil, err
	}
	return hrp, converted, nil
}
