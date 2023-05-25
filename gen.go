//go:generate abigen --abi=./mainnet/abi/redistribution.abi --pkg=redistribution --out=./mainnet/redistribution/redistribution.go
//go:generate abigen --abi=./mainnet/abi/staking.abi --pkg=staking --out=./mainnet/staking/staking.go
//go:generate abigen --abi=./mainnet/abi/postageStamp.abi --pkg=postagestamp --out=./mainnet/postagestamp/postagestamp.go
//go:generate abigen --abi=./mainnet/abi/priceOracle.abi --pkg=priceoracle --out=./mainnet/priceoracle/priceoracle.go

//go:generate abigen --abi=./testnet/abi/redistribution.abi --pkg=redistribution --out=./testnet/redistribution/redistribution.go
//go:generate abigen --abi=./testnet/abi/staking.abi --pkg=staking --out=./testnet/staking/staking.go
//go:generate abigen --abi=./testnet/abi/postageStamp.abi --pkg=postagestamp --out=./testnet/postagestamp/postagestamp.go
//go:generate abigen --abi=./testnet/abi/priceOracle.abi --pkg=priceoracle --out=./testnet/priceoracle/priceoracle.go

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/ethersphere/go-storage-incentives-abi/abi"
)

func main() {
	mustWrite("mainnet", "staking", abi.MainnetStakingABI)
	mustWrite("mainnet", "redistribution", abi.MainnetRedistributionABI)
	mustWrite("mainnet", "postageStamp", abi.MainnetPostageStampStampABI)
	mustWrite("mainnet", "priceOracle", abi.MainnetPriceOracleABI)

	mustWrite("testnet", "staking", abi.TestnetStakingABI)
	mustWrite("testnet", "redistribution", abi.TestnetRedistributionABI)
	mustWrite("testnet", "postageStamp", abi.TestnetPostageStampABI)
	mustWrite("testnet", "priceOracle", abi.MainnetPriceOracleABI)
}

func mustWrite(target, fileName, contents string) {
	err := os.WriteFile(fmt.Sprintf("%s/abi/%s.abi", target, fileName), []byte(contents), 0644)
	if err != nil {
		log.Fatal(err)
	}
}
