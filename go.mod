module u-sdk

go 1.20

require (
	github.com/btcsuite/btcutil v1.0.3-0.20201208143702-a53e38424cce
	github.com/ethereum/go-ethereum v1.11.2
	github.com/irisnet/core-sdk-go v0.2.1-0.20230606070930-9082c23b90fc
	github.com/stretchr/testify v1.8.1
	github.com/tyler-smith/go-bip39 v1.1.0
)

require (
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

require (
	github.com/ChainSafe/go-schnorrkel v0.0.0-20200405005733-88cbf1b4c40d // indirect
	github.com/btcsuite/btcd v0.22.0-beta // indirect
	github.com/btcsuite/btcd/btcec/v2 v2.2.0 // indirect
	github.com/cosmos/cosmos-sdk v0.44.3 // indirect
	github.com/cosmos/go-bip39 v1.0.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/decred/dcrd/dcrec/secp256k1/v4 v4.0.1 // indirect
	github.com/go-kit/kit v0.10.0 // indirect
	github.com/go-logfmt/logfmt v0.5.0 // indirect
	github.com/gogo/protobuf v1.3.3 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	github.com/gtank/merlin v0.1.1 // indirect
	github.com/gtank/ristretto255 v0.1.2 // indirect
	github.com/miguelmota/go-ethereum-hdwallet v0.1.1
	github.com/mimoo/StrobeGo v0.0.0-20181016162300-f8f6d4d2b643 // indirect
	github.com/petermattis/goid v0.0.0-20220712135657-ac599d9cba15 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/sasha-s/go-deadlock v0.2.0 // indirect
	github.com/spf13/cobra v1.8.0
	github.com/tendermint/crypto v0.0.0-20191022145703-50d29ede1e15 // indirect
	github.com/tendermint/go-amino v0.16.0 // indirect
	github.com/tendermint/tendermint v0.34.14 // indirect
	github.com/tjfoc/gmsm v1.4.0 // indirect
	golang.org/x/crypto v0.12.0 // indirect
	golang.org/x/net v0.10.0 // indirect
	golang.org/x/sys v0.11.0 // indirect
	golang.org/x/text v0.12.0 // indirect
	google.golang.org/genproto v0.0.0-20230110181048-76db0878b65f // indirect
	google.golang.org/grpc v1.53.0 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

replace (
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	//github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1
	//github.com/irisnet/core-sdk-go v0.0.0-20220905061348-af15d19e6045 => github.com/irisnet/core-sdk-go v0.0.0-20230316082027-6974c08df472
	github.com/irisnet/core-sdk-go v0.0.0-20220905061348-af15d19e6045 => github.com/irisnet/core-sdk-go v0.2.1-0.20230606070930-9082c23b90fc
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.34.1-irita-210113
)
