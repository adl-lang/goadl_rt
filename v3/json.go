package goadl

import (
	"bytes"
	"reflect"
	"strings"
)

type EncodeState struct {
	bytes.Buffer // accumulated output
}

type EncoderFunc func(e *EncodeState, v reflect.Value) error

type DecodeState struct {
	V    reflect.Value
	Path ctxPath
}

type ctxPath []string

func (cp ctxPath) String() string {
	return "[" + strings.Join(cp, ",") + "]"
}

type DecodeFunc func(ds *DecodeState, v any) error
