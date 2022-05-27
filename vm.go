package vm

import (
	"bytes"
	"errors"
	"fmt"

	"github.com/spacemeshos/go-scale"
)

type VM struct {
	storage Storage
}

type Request struct {
	vm *VM

	raw     []byte
	ctx     *Context
	decoder *scale.Decoder
	args    any
}

func (r *Request) Parse() (*Header, error) {
	header, ctx, args, err := r.vm.parse(r.decoder)
	if err != nil {
		return nil, err
	}
	r.ctx = ctx
	r.args = args
	return header, nil
}

func (r *Request) Verify() bool {
	return r.vm.verify(r.ctx, r.raw)
}

func (vm *VM) Validation(raw []byte) *Request {
	return &Request{
		decoder: scale.NewDecoder(bytes.NewReader(raw)),
		raw:     raw,
	}
}

func (vm *VM) parse(decoder *scale.Decoder) (*Header, *Context, any, error) {
	version, _, err := scale.DecodeCompact8(decoder)
	if err != nil {
		return nil, nil, nil, err
	}
	if version != 1 {
		return nil, nil, nil, fmt.Errorf("unsupported version %d", version)
	}

	var principal scale.Address
	if _, err := principal.DecodeScale(decoder); err != nil {
		return nil, nil, nil, err
	}
	method, _, err := scale.DecodeCompact8(decoder)
	if err != nil {
		return nil, nil, nil, err
	}
	// fetch whole state in a single scan
	state := vm.storage.Get(principal)
	if method == 0 {
		var template scale.Address
		if _, err := template.DecodeScale(decoder); err != nil {
			return nil, nil, nil, err
		}
		state.Template = &template
	}
	handler := templates[*state.Template]
	if handler == nil {
		return nil, nil, nil, errors.New("unknown template")
	}
	ctx := &Context{
		handler:   handler,
		State:     state,
		principal: principal,
		method:    method,
	}
	header, args := handler.Parse(ctx, method, decoder)
	ctx.Template = handler.Load(ctx, method, &header)
	return &header, ctx, args, nil
}

func (vm *VM) verify(ctx *Context, raw []byte) bool {
	return ctx.Template.Verify(ctx, raw)
}

// Apply transaction. Returns true if transaction run out of gas in the validation phase.
func (vm *VM) Apply(txs [][]byte) List[[]byte] {
	var (
		changes List[*AccountState]
		rd      bytes.Reader
		decoder = scale.NewDecoder(&rd)
		failed  List[[]byte]
	)
	for _, tx := range txs {
		rd.Reset(tx)
		_, ctx, args, err := vm.parse(decoder)
		if err != nil {
			failed.Add(tx)
		}
		// TODO skip verification but consume cost
		ctx.Consume(100)
		ctx.handler.Exec(ctx, ctx.method, args)
		changes.Add(ctx.State)
	}
	iterator := changes.Iterate()
	for {
		state, next := iterator.Next()
		if !next {
			break
		}
		vm.storage.Put(Address{}, state)
	}
	return failed
}
