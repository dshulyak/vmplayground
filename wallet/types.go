package wallet

import "github.com/spacemeshos/vm"

//go:generate scalegen -pkg wallet -file types_scale.go -types Arguments,SpawnArguments,SpendPayload,Nonce,SpawnPayload -imports github.com/spacemeshos/vm/wallet

type SpawnArguments struct {
	PublicKey vm.PublicKey
}

type Arguments struct {
	Destination vm.Address
	Amount      uint64
}

type SpendPayload struct {
	Arguments Arguments
	Nonce     Nonce
	GasPrice  uint32
}

type Nonce struct {
	Counter  uint64
	Bitfield uint8
}

type SpawnPayload struct {
	Arguments SpawnArguments
	GasPrice  uint32
}
