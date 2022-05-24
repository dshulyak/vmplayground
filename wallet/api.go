package wallet

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/vm"
)

func Parse(ctx *vm.Context, method uint8, decoder *scale.Decoder) (payload vm.Header) {
	ctx.Consume(10)
	switch method {
	case 0:
		var p SpawnPayload
		if _, err := p.DecodeScale(decoder); err != nil {
			ctx.Fail(errors.New("invalid tx"))
		}
		payload.Arguments = p.Arguments
		payload.GasPrice = uint64(p.GasPrice)
	case 1:
		var p SpendPayload
		if _, err := p.DecodeScale(decoder); err != nil {
			ctx.Fail(errors.New("invalid tx"))
		}
		payload.Arguments = p.Arguments
		payload.GasPrice = uint64(p.GasPrice)
		payload.Nonce.Counter = p.Nonce.Counter
		payload.Nonce.Bitfield = p.Nonce.Bitfield
	}
	return payload
}

func Load(ctx *vm.Context, method uint8, header *vm.Header) any {
	if method == 0 {
		if ctx.State.Template != nil {
			ctx.Fail(errors.New("account already spawned"))
		}
		return New(ctx, header.Arguments.(SpawnArguments))
	}
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

func Exec(ctx *vm.Context, method uint8, header *vm.Header) {
	switch method {
	case 0:
		ctx.Consume(100)
	case 1:
		ctx.Consume(50)
		ctx.Template.(*Wallet).Spend(ctx, header.Arguments.(*Arguments))
	default:
		ctx.Fail(vm.ErrUnknownMethodSelector)
	}
}
