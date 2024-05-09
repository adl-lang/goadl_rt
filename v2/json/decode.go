package json

import (
	"io"

	goadl "github.com/adl-lang/goadl_rt/v2"
)

type Decoder struct {
	r io.Reader
}

func NewDecoder[T any](
	r io.Reader,
	texpr goadl.ATypeExpr[T],
) *Decoder {
	return &Decoder{r: r}
}
