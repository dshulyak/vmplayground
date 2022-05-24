package vm

type Header struct {
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
