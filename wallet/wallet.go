package wallet

import (
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"

	"github.com/spacemeshos/vm"
)

func New(ctx *vm.Context, args SpawnArguments) *Wallet {
	return &Wallet{PublicKey: args.PublicKey}
}

//go:generate scalegen -pkg wallet -file wallet_scale.go -types Wallet -imports github.com/spacemeshos/vm/wallet

type Wallet struct {
	PublicKey vm.PublicKey
}

func (s *Wallet) MaxSpend(method uint8, payload *vm.Header) uint64 {
	switch method {
	case 0:
		return 0
	case 1:
		return payload.Arguments.(*Arguments).Amount
	default:
		panic("unreachable")
	}
}

func (s *Wallet) Verify(tx []byte) bool {
	if len(tx) < 64 {
		return false
	}
	return ed25519.Verify(ed25519.PublicKey(s.PublicKey[:]), tx[:len(tx)-64], tx[len(tx)-64:])
}

func (s *Wallet) Spend(ctx *vm.Context, args *Arguments) {
	ctx.Transfer(args.Destination, args.Amount)
}
