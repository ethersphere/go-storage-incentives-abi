.DEFAULT_GOAL := gen

.PHONY: clean

gen:
	@mkdir -p mainnet/abi mainnet/redistribution mainnet/staking mainnet/postagestamp mainnet/priceoracle
	@mkdir -p testnet/abi testnet/redistribution testnet/staking testnet/postagestamp testnet/priceoracle
	go run .
	go generate ./...

clean:
	@rm -rf ./testnet ./mainnet