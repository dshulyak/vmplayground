// Code generated by github.com/spacemeshos/go-scale/gen. DO NOT EDIT.

package vm

import (
	"github.com/spacemeshos/go-scale"
)

func (t *Nonce) EncodeScale(enc *scale.Encoder) (total int, err error) {
	// field Counter (0)
	if n, err := scale.EncodeCompact64(enc, t.Counter); err != nil {
		return total, err
	} else {
		total += n
	}

	// field Bitfield (1)
	if n, err := scale.EncodeCompact8(enc, t.Bitfield); err != nil {
		return total, err
	} else {
		total += n
	}

	return total, nil
}

func (t *Nonce) DecodeScale(dec *scale.Decoder) (total int, err error) {
	// field Counter (0)
	if field, n, err := scale.DecodeCompact64(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Counter = field
	}

	// field Bitfield (1)
	if field, n, err := scale.DecodeCompact8(dec); err != nil {
		return total, err
	} else {
		total += n
		t.Bitfield = field
	}

	return total, nil
}
