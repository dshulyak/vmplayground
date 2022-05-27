package multisig

import "github.com/spacemeshos/go-scale"

type Signature struct {
	Conf uint8
	Sigs [2]scale.Signature
}
