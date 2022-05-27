package vm

import (
	"bytes"
	"crypto/sha256"
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
	storage  Storage
	handler  *TemplateAPI
	Template Template

	templateAddress Address
	header          Header
	args            scale.Encodable

	State     *AccountState
	principal Address
	method    uint8

	spawned struct {
		address scale.Address
		state   *AccountState
	}
	consumed uint64
	price    uint64
	funds    uint64
}

func (c *Context) Spawn() {
	hasher := sha256.New()
	encoder := scale.NewEncoder(hasher)

	c.templateAddress.EncodeScale(encoder)
	c.header.Nonce.EncodeScale(encoder)
	c.args.EncodeScale(encoder)
	hasher.Sum(c.spawned.address[:])

	if c.spawned.address == c.principal {
		c.spawned.state = c.State
	} else {
		state := c.storage.Get(c.spawned.address)
		c.spawned.state = state
	}
	if c.spawned.state.Template != nil {
		c.Fail(errors.New("already spawned"))
	}
	buf := bytes.NewBuffer(nil)
	encoder = scale.NewEncoder(buf)
	c.Template.EncodeScale(encoder)

	c.spawned.state.Template = &c.templateAddress
	c.spawned.state.State = buf.Bytes()
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
