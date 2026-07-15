package rwa

//go:generate abigen -abi Order.abi.json -out order.go -pkg rwa -type Order
//go:generate abigen -abi CBJToken.abi.json -out cbj_token.go -pkg rwa -type CBJToken
//go:generate abigen -abi CBJGateway.abi.json -out cbj_gateway.go -pkg rwa -type CBJGateway
