
# 命令说明
```
./build/cbac -h
公私钥生成和转换

Usage:
  cbac [flags]
  cbac [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  evm2iaa     evm的address转换为原生的iaa地址
  evmpri2iaa  根据EVM的私钥生成address
  genevm      生成EVM的 助记词，私钥，公钥，地址
  help        Help about any command
  iaa2evm     iaa的address转换为evm地址

Flags:
  -h, --help   help for cbac

Use "cbac [command] --help" for more information about a command.
```

./cbac genevm
eth_secp256k1
nmemonic(助记词):                smile clown ghost usage movie burger deal behave milk relax across cross unit true occur begin cotton casino category tomato enhance pink kiwi utility
privateKey(私钥):                02fcf0c81ecc8c1945b2e98879c45788ce6d38ef6ee735d4c9a6927085560a48
publicKey(公钥):                 e312f5354530b5adf42ebeeda735ff314c1f6d9152bca66ec816101bea351030939b5999177d57ed00c919cc58f24275766ef11beedab04078751272f7cfd897
address(EVM地址):                0xDEeCD112c111c8c2201Eb3C701d7C1D5D806E931
address(原生地址):               iaa1mmkdzykpz8yvygq7k0rsr47p6hvqd6f3lqgnvy
privateKeyIAA(原生私钥):         0xfcd2efcc2002fcf0c81ecc8c1945b2e98879c45788ce6d38ef6ee735d4c9a6927085560a48
publicKeyIAA(原生公钥):          0xf3b3cd032103e312f5354530b5adf42ebeeda735ff314c1f6d9152bca66ec816101bea351030

这个是随机生成的，每次都不同，请记住所有的参数，一旦忘记，就无法操作账号。