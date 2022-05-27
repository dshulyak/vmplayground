package wallet

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/vm"
)

func Parse(ctx *vm.Context, method uint8, decoder *scale.Decoder) (header vm.Header, args scale.Encodable) {
	// TODO rethink cost approach
	ctx.Consume(10)
	switch method {
	case 0:
		var p SpawnPayload
		if _, err := p.DecodeScale(decoder); err != nil {
			ctx.Fail(errors.New("invalid tx"))
		}
		args = &p.Arguments
		header.GasPrice = uint64(p.GasPrice)
	case 1:
		var p SpendPayload
		if _, err := p.DecodeScale(decoder); err != nil {
			ctx.Fail(errors.New("invalid tx"))
		}
		args = &p.Arguments
		header.GasPrice = uint64(p.GasPrice)
		header.Nonce.Counter = p.Nonce.Counter
		header.Nonce.Bitfield = p.Nonce.Bitfield
	}
	return header, args
}

func Load(ctx *vm.Context, method uint8, args any) vm.Template {
	if method == 0 {
		return New(args.(SpawnArguments))
	}
	// TODO i dont like that it needs to initialize decoder here
	// what can be done about it?
	decoder := scale.NewDecoder(bytes.NewReader(ctx.State.State))
	var wallet Wallet
	if _, err := wallet.DecodeScale(decoder); err != nil {
		ctx.Fail(fmt.Errorf("internal error %w", err))
	}
	return &wallet
}

func Exec(ctx *vm.Context, method uint8, args any) {
	switch method {
	case 0:
		ctx.Consume(100)
		ctx.Spawn()
	case 1:
		ctx.Consume(50)
		ctx.Template.(*Wallet).Spend(ctx, args.(*Arguments))
	default:
		// TODO change it to propagate errors without throws
		ctx.Fail(vm.ErrUnknownMethodSelector)
	}
}
