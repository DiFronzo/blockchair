<h1 align="center">
  <img src="./misc/BC_api.svg" width="700px"/><br/>
   Blockchair API
</h1>
<p align="center">Client for the blockchair.com API using GO.

<p align="center"><a href="https://github.com/DiFronzo/blockchair/releases" target="_blank"><img src="https://img.shields.io/badge/version-v0.1.0-blue?style=for-the-badge&logo=none" alt="BC-API version" /></a>&nbsp;<a href="https://golang.org/" target="_blank"><img src="https://img.shields.io/badge/GO-1.17+-00ADD8?style=for-the-badge&logo=GO" alt="go version" /></a>&nbsp;<img src="https://img.shields.io/badge/license-MIT-red?style=for-the-badge&logo=none" alt="license" />&nbsp;<img alt="code size" src="https://img.shields.io/github/languages/code-size/difronzo/blockchair?style=for-the-badge&logo=none">&nbsp;<a href="https://goreportcard.com/report/github.com/DiFronzo/blockchair" target="_blank"><img src="https://goreportcard.com/badge/github.com/DiFronzo/blockchair?style=for-the-badge&logo=none" alt="GO report" />&nbsp;<a href="https://pkg.go.dev/github.com/DiFronzo/blockchair" target="_blank"><img src="https://img.shields.io/badge/GoDoc-reference-blue?style=for-the-badge&logo=go" alt="GoDoc" /></a>&nbsp;<a href="https://github.com/DiFronzo/blockchair/actions" target="_blank"><img alt="GitHub Workflow Status" src="https://img.shields.io/github/workflow/status/difronzo/blockchair/Tests?logo=github&style=for-the-badge"/></a></p>

## ⚡️ Quick start

First of all, [download](https://golang.org/) and install **GO**. Version `1.17` or higher is required.

Verify that the installation was successful by running the following command that should return the version number for GO.

```bash
go version
```

To quickly start using the module run the following command for installation.

```bash
go install github.com/DiFronzo/blockchair@latest
```

That's all you need to know to start! 🎉

## ⚙️ Usage & Options

### `Get address for Bitcoin-like crypto`
The function `GetAddress` is used to find information regarding a specific address. The function takes two arguments, the type of crypto (string) and the address (string).
```go
package main

import (
	"fmt"
	"log"

	"github.com/DiFronzo/blockchair"
)

func main() {
	c, _ := blockchair.New()
	resp, err := c.GetAddress("bitcoin", "34xp4vRoCGJym3xR7yCVPFHoCNxv4Twseo")
	if err != nil {
		log.Fatalln(err)
	}
	for i := range resp.Data{
		fmt.Println(resp.Data[i].Address.Type)
    }
}
```

### `Get addresses for Bitcoin-like crypto`
The function `GetAddresses` is used to find information regarding multiple addresses. The function takes two arguments, the type of crypto (string) and the addresses ([]string).
```go
package main

import (
	"fmt"
	"log"

	"github.com/DiFronzo/blockchair"
)

func main() {
	c, _ := blockchair.New()
	resp, err := c.GetAddresses("bitcoin", []string{"34xp4vRoCGJym3xR7yCVPFHoCNxv4Twseo","bc1qgdjqv0av3q56jvd82tkdjpy7gdp9ut8tlqmgrpmv24sq90ecnvqqjwvw97"})
	if err != nil {
		log.Fatalln(err)
	}
	for i := range resp.Data.Addresses {
		fmt.Println(resp.Data.Addresses[i].Type)
    }
}
```
### `Get address for Ethereum`
The function `GetAddressEth` is used to find information regarding a specific Ethereum address. Identical to [Get address for Bitcoin-like crypto](https://github.com/DiFronzo/blockchair#get-address-for-bitcoin-like-crypto) just with the function `GetAddressEth`.

### `Get addresses for Ethereum`
The function `GetAddressesEth` is used to find information regarding multiple Ethereum addresses. Identical to [Get addresses for Bitcoin-like crypto](https://github.com/DiFronzo/blockchair#get-addresses-for-bitcoin-like-crypto) just with the function `GetAddressesEth`.

### 🐳 Docker-way to quick start

UNDER CONSTRUCTION.

## ⭐️ Project assistance

If you want to say **thank you** or/and support active development of `blockchair`:

- Add a [GitHub Star](https://github.com/DiFronzo/blockchair) to the project.

## ⚠️ License
`Blockchair` is free and open-source software licensed under the [MIT](https://github.com/DiFronzo/blockchair/blob/main/LICENSE). This is not an offical release from [Blockchair](https://github.com/Blockchair). Use on your own risk.
