package vm

import (
	"errors"

	"github.com/spacemeshos/go-scale"
)

var (
	ErrUnknownMethodSelector = errors.New("unknown method selector")
	ErrOutOfGas              = errors.New("out of gas")
)

type (
	PublicKey = scale.PublicKey
	Address   = scale.Address
)

type Host struct {
	storage Storage

	touched List[Address]
	state   map[Address]AccountState
}

type AccountState struct {
	// State is decoded into precompile as is
	State    []byte
	Balance  uint64
	Template *Address // only available for spawned account
}

type Storage interface {
	Get(Address) *AccountState
	Put(Address, *AccountState)
}

type Context struct {
	handler   *TemplateAPI
	Template  any
	State     *AccountState
	principal Address
	method    uint8

	consumed uint64
	price    uint64
	funds    uint64
}

func (c *Context) Transfer(to Address, value uint64) {}

func (c *Context) Consume(gas uint64) {
	c.consumed += gas
	if c.consumed*c.price+c.funds > c.State.Balance {
		c.Fail(errors.New("out of funds"))
	}
}

func (c *Context) Fail(err error) {
	//
	panic(err)
}
