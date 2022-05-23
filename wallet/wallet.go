package wallet

import (
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"

	"github.com/spacemeshos/vm"
	"github.com/spacemeshos/vm/generic"
)

func New(ctx *vm.Context, args SpawnArguments) Single {
	return Single{PublicKey: args.PublicKey}
}

//go:generate scalegen -pkg wallet -file wallet_scale.go -types Single -imports github.com/spacemeshos/vm/wallet

type Single struct {
	PublicKey vm.PublicKey
}

func (s *Single) MaxSpend(method uint8, payload *generic.Payload) uint64 {
	switch method {
	case 0:
		return 0
	case 1:
		return payload.Arguments.(*Arguments).Amount
	default:
		panic("unreachable")
	}
}

func (s *Single) Verify(tx []byte) bool {
	if len(tx) < 64 {
		return false
	}
	return ed25519.Verify(ed25519.PublicKey(s.PublicKey[:]), tx[:len(tx)-64], tx[len(tx)-64:])
}

func (s *Single) Spend(ctx *vm.Context, args *Arguments) {
	ctx.Transfer(args.Destination, args.Amount)
}
