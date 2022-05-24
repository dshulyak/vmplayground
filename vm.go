package vm

import (
	"errors"

	"github.com/spacemeshos/go-scale"
)

type VM struct {
	storage Storage
}

func (vm *VM) Verify(raw []byte, decoder *scale.Decoder) (*Header, error) {
	var principal scale.Address
	if _, err := principal.DecodeScale(decoder); err != nil {
		return nil, err
	}
	method, _, err := scale.DecodeCompact8(decoder)
	if err != nil {
		return nil, err
	}
	// fetch whole state in a single scan
	state := vm.storage.Get(principal)
	if method == 0 {
		var template scale.Address
		if _, err := template.DecodeScale(decoder); err != nil {
			return nil, err
		}
		state.Template = &template
	}
	handler := templates[*state.Template]
	if handler == nil {
		return nil, errors.New("unknown template")
	}
	ctx := &Context{
		State:     state,
		principal: principal,
	}
	header := handler.Parse(ctx, method, decoder)
	spawned := handler.Load(ctx, method, &header)
	ctx.Template = spawned
	if !handler.Verify(ctx, raw) {
		return nil, errors.New("verification failed")
	}
	return &header, nil
}

// Apply transaction. Returns true if transaction run out of gas in the validation phase.
func (vm *VM) Apply(decoder *scale.Decoder) bool {
	return true
}
