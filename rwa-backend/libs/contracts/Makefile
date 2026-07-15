export-abi:
	cd ../../../rwa-contract && forge inspect OrderContract abi --json > build/Order.abi.json
	cd ../../../rwa-contract && forge inspect CBJToken abi --json > build/CBJToken.abi.json
	cd ../../../rwa-contract && forge inspect CBJGateway abi --json > build/CBJGateway.abi.json
	cp ../../../rwa-contract/build/Order.abi.json rwa/Order.abi.json
	cp ../../../rwa-contract/build/CBJToken.abi.json rwa/CBJToken.abi.json
	cp ../../../rwa-contract/build/CBJGateway.abi.json rwa/CBJGateway.abi.json

all: export-abi generate

.PHONY: generate export-abi
generate:
	cd rwa && go generate ./...