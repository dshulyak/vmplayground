package generic

import (
	"bytes"

	"github.com/spacemeshos/go-scale"
	"github.com/spacemeshos/vm"
)

type Payload struct {
	Arguments any
	Nonce     struct {
		Counter  uint64
		Bitfield uint8
	}
	Layer struct {
		Min, Max uint32
	}
	MaxGas   uint64
	GasPrice uint64
	MaxSpend uint64
}

func Spawn[T1, T2 any, H1 scale.TypeHelper[T1], H2 scale.TypeHelper[T2]](ctx *vm.Context, spawner func(*vm.Context, T1) T2, decoder *scale.Decoder) {
	var args T1
	if _, err := H1(&args).DecodeScale(decoder); err != nil {
		panic(err)
	}
	state := spawner(ctx, args)
	encoder := scale.NewEncoder(bytes.NewBuffer(nil))
	if _, err := H2(&state).EncodeScale(encoder); err != nil {
		panic(err)
	}
}
