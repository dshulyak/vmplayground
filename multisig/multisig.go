package multisig

import "github.com/spacemeshos/go-scale"

const required = 2

type MultiSig struct {
	Keys [3]scale.PublicKey
}
