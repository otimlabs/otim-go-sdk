.PHONY: update-otim-protocol eth-devtools

otim-protocol:
	git clone git@github.com:otimlabs/otim-protocol.git

update-otim-protocol:
	cd otim-protocol && git pull

go-ethereum:
	git clone git@github.com:ethereum/go-ethereum.git

eth-devtools: go-ethereum
	cd go-ethereum && make && make devtools

foundry:
	curl -L https://foundry.paradigm.xyz | bash

foundryup:
	foundryup

otim-protocol-abi:
	cd otim-protocol && forge soldeer update && forge build --extra-output-files abi

go-abi: otim-protocol-abi
	./scripts/gen_abi.sh