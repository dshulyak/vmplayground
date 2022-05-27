package vm

type Header struct {
	Nonce Nonce
	Layer struct {
		Min, Max uint32
	}
	MaxGas   uint64
	GasPrice uint64
	MaxSpend uint64
}

//go:generate scalegen -pkg vm -file header_scale.go -types Nonce -imports github.com/spacemeshos/vm

type Nonce struct {
	Counter  uint64
	Bitfield uint8
}
