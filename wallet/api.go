package wallet

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/vm"
)

func Parse(ctx *vm.Context, method uint8, decoder *scale.Decoder) (payload vm.Header, args any) {
	// TODO rethink cost approach
	ctx.Consume(10)
	switch method {
	case 0:
		var p SpawnPayload
		if _, err := p.DecodeScale(decoder); err != nil {
			ctx.Fail(errors.New("invalid tx"))
		}
		args = p.Arguments
		payload.GasPrice = uint64(p.GasPrice)
	case 1:
		var p SpendPayload
		if _, err := p.DecodeScale(decoder); err != nil {
			ctx.Fail(errors.New("invalid tx"))
		}
		args = p.Arguments
		payload.GasPrice = uint64(p.GasPrice)
		payload.Nonce.Counter = p.Nonce.Counter
		payload.Nonce.Bitfield = p.Nonce.Bitfield
	}
	return payload, args
}

func Load(ctx *vm.Context, method uint8, args any) any {
	if method == 0 {
		if ctx.State.Template != nil {
			ctx.Fail(errors.New("account already spawned"))
		}
		return New(ctx, args.(SpawnArguments))
	}
	// TODO i dont like that it needs to initialized decoder here
	// what can be done about it?
	decoder := scale.NewDecoder(bytes.NewReader(ctx.State.State))
	var wallet Wallet
	if _, err := wallet.DecodeScale(decoder); err != nil {
		ctx.Fail(fmt.Errorf("internval error %w", err))
	}
	return &wallet
}

func Verify(ctx *vm.Context, raw []byte) bool {
	ctx.Consume(50)
	return ctx.Template.(*Wallet).Verify(raw)
}

func Exec(ctx *vm.Context, method uint8, args any) {
	switch method {
	case 0:
		ctx.Consume(100)
	case 1:
		ctx.Consume(50)
		ctx.Template.(*Wallet).Spend(ctx, args.(*Arguments))
	default:
		// TODO change it to propagate errors without throws
		ctx.Fail(vm.ErrUnknownMethodSelector)
	}
}
