package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"u-sdk/blockchain/eth"
	"u-sdk/blockchain/irita"
)

var rootCmd = &cobra.Command{
	Use:   "cbac",
	Short: "公私钥生成和转换",
	Long:  `公私钥生成和转换`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("公私钥生成")
	},
}

var genevm = &cobra.Command{
	Use:   "genevm",
	Short: "生成EVM的 助记词，私钥，公钥，地址",
	Long:  `生成EVM的 助记词，私钥，公钥，地址`,
	Run: func(cmd *cobra.Command, args []string) {
		var acc = eth.Account{}
		nmemonic, privateKeyHex, publicKeyHex, address, err := acc.HDWalletGenerate()
		if err != nil {
			fmt.Println(err)
		}
		addressIAA, err := irita.HexToBech32(address)
		if err != nil {
			fmt.Println(err)
		}
		privateKeyIAA, publicKeyIAA, err := acc.EvmPrivateKeyToIAA(privateKeyHex)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("nmemonic(助记词):\t\t", nmemonic)
		fmt.Println("privateKey(私钥):\t\t", privateKeyHex)
		fmt.Println("publicKey(公钥):\t\t", publicKeyHex)
		fmt.Println("address(EVM地址):\t\t", address)
		fmt.Println("address(原生地址):\t\t", addressIAA)
		fmt.Println("privateKeyIAA(原生私钥):\t", privateKeyIAA)
		fmt.Println("publicKeyIAA(原生公钥):\t\t", publicKeyIAA)
	},
}

var evm2iaa = &cobra.Command{
	Use:   "evm2iaa",
	Short: "evm的address转换为原生的iaa地址",
	Long:  `evm的address转换为原生的iaa地址`,
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		addressIAA, err := irita.HexToBech32(address)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(addressIAA)
	},
}
var iaa2evm = &cobra.Command{
	Use:   "iaa2evm",
	Short: "iaa的address转换为evm地址",
	Long:  `iaa的address转换为evm地址`,
	Run: func(cmd *cobra.Command, args []string) {
		address, _ := cmd.Flags().GetString("address")
		addressEVM, err := irita.Bech32ToHex(address, "iaa")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(addressEVM)
	},
}
var evmpri2iaa = &cobra.Command{
	Use:   "evmpri2iaa",
	Short: "根据EVM的私钥生成address",
	Long:  `根据EVM的私钥生成address`,
	Run: func(cmd *cobra.Command, args []string) {
		privateKeyEVM, _ := cmd.Flags().GetString("privatekey")
		var acc = eth.Account{}
		privateKeyIAA, publicKeyIAA, err := acc.EvmPrivateKeyToIAA(privateKeyEVM)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("privateKeyIAA:\t", privateKeyIAA)
		fmt.Println("publicKeyIAA:\t", publicKeyIAA)
	},
}

func init() {
	rootCmd.AddCommand(genevm, evm2iaa, iaa2evm, evmpri2iaa)
	evm2iaa.Flags().StringP("address", "a", "", "EVM的0x地址")
	iaa2evm.Flags().StringP("address", "a", "", "IAA的0x地址")
	evmpri2iaa.Flags().StringP("privatekey", "k", "", "EVM私钥")

}
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
