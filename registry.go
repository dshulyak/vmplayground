package vm

import "github.com/spacemeshos/go-scale"

var (
	templates = map[Address]*TemplateAPI{}
)

type TemplateAPI struct {
	Parse func(*Context, uint8, *scale.Decoder) (Header, scale.Encodable)
	Load  func(*Context, uint8, any) Template
	Exec  func(*Context, uint8, any)
}

type Template interface {
	scale.Encodable
	Verify(*Context, []byte) bool
}
