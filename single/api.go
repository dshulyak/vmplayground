package single

import (
	"errors"

	"github.com/spacemeshos/go-scale"

	"github.com/spacemeshos/vm"
	"github.com/spacemeshos/vm/generic"
)

func Parse(ctx *vm.Context, method uint8, decoder *scale.Decoder) (payload generic.Payload) {
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

func Load(ctx *vm.Context, method uint8, payload *generic.Payload) any {

	return nil
}

func Verify(ctx *vm.Context, raw []byte) bool {
	ctx.Consume(50)
	return ctx.Template.(*Single).Verify(raw)
}

func Exec(ctx *vm.Context, method uint8, payload *generic.Payload) {
	switch method {
	case 0:
		ctx.Consume(100)
	case 1:
		ctx.Consume(50)
		ctx.Template.(*Single).Spend(ctx, payload.Arguments.(*Arguments))
	default:
		ctx.Fail(vm.ErrUnknownMethodSelector)
	}
}
