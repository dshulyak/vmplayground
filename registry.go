package vm

import "github.com/spacemeshos/go-scale"

var (
	templates = map[Address]*TemplateAPI{}
)

type TemplateAPI struct {
	Parse  func(*Context, uint8, *scale.Decoder) Header
	Load   func(*Context, uint8, *Header) any
	Verify func(*Context, []byte) bool
	Exec   func(*Context, uint8, *Header)
}
