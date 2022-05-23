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
	Template Address
}

type Storage interface {
	Get(Address) *AccountState
	Put(Address, *AccountState)
}

type Context struct {
	Template  any
	State     *AccountState
	principal Address
	host      *Host

	consumed uint64
}

func (c *Context) Load(addr Address) {
	c.State = c.host.storage.Get(addr)
}

func (c *Context) Transfer(to Address, value uint64) {}

func (c *Context) Consume(gas uint64) {
	c.consumed += gas
}

func (c *Context) Fail(err error) {
	panic(err)
}
