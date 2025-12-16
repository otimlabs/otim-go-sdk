.PHONY: update-otim-protocol eth-devtools

otim-protocol:
	git clone git@github.com:otimlabs/otim-protocol.git

update-otim-protocol:
	cd otim-protocol && git pull

eth-devtools:
	go install github.com/ethereum/go-ethereum/cmd/abigen@v1.16.7

foundry:
	curl -L https://foundry.paradigm.xyz | bash

foundryup:
	foundryup

otim-protocol-abi:
	cd otim-protocol && forge soldeer update && forge build --extra-output-files abi

install-otim-protocol-deps: foundry foundryup

go-abi-gen: eth-devtools otim-protocol-abi
	./scripts/gen_abi.sh