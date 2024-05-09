package json

import (
	"bytes"
	"encoding"
	"encoding/base64"
	gojson "encoding/json"
	"fmt"
	"io"
	"math"
	"reflect"
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"

	. "github.com/adl-lang/goadl_rt/v2"
	goadl "github.com/adl-lang/goadl_rt/v2"
)

type Encoder[T any] struct {
	w       io.Writer
	binding encoderFunc
}

func NewEncoder[T any](
	w io.Writer,
	texpr goadl.ATypeExpr[T],
	dres goadl.Resolver,
) *Encoder[T] {
	binding := buildEncodeBinding(dres, texpr.Value, make(boundEncodeTypeParams))
	return &Encoder[T]{
		w:       w,
		binding: binding,
	}
}

func (enc *Encoder[T]) Encode(v T) error {
	es := &encodeState{}
	enc.binding(es, reflect.ValueOf(v))
	b := es.Bytes()
	_, err := enc.w.Write(b)
	return err
}

type encodeState struct {
	bytes.Buffer // accumulated output
}

type encoderFunc func(e *encodeState, v reflect.Value)

// var encoderCache sync.Map // map[reflect.Type]encoderFunc

// func typeEncoder(t reflect.Type) encoderFunc {
// 	if fi, ok := encoderCache.Load(t); ok {
// 		return fi.(encoderFunc)
// 	}

// 	// To deal with recursive types, populate the map with an
// 	// indirect func before we build it. This type waits on the
// 	// real func (f) to be ready and then calls it. This indirect
// 	// func is only used for recursive types.
// 	var (
// 		wg sync.WaitGroup
// 		f  encoderFunc
// 	)
// 	wg.Add(1)
// 	fi, loaded := encoderCache.LoadOrStore(t, encoderFunc(func(e *encodeState, v reflect.Value) {
// 		wg.Wait()
// 		f(e, v)
// 	}))
// 	if loaded {
// 		return fi.(encoderFunc)
// 	}
// 	// Compute the real encoder and replace the indirect func with it.
// 	f = newTypeEncoder(t, true)
// 	wg.Done()
// 	encoderCache.Store(t, f)
// 	return f
// }

type boundEncodeTypeParams map[string]encoderFunc

func buildEncodeBinding(
	dres Resolver,
	texpr TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	return Ok(HandleTypeRef[encoderFunc](
		texpr.TypeRef.Branch,
		func(trb TypeRefBranch_Primitive) (encoderFunc, error) {
			return primitiveEncodeBinding(dres, string(trb), texpr.Parameters, boundTypeParams), nil
		},
		func(trb TypeRefBranch_TypeParam) (encoderFunc, error) {
			return boundTypeParams[string(trb)], nil
		},
		func(trb TypeRefBranch_Reference) (encoderFunc, error) {
			ast := dres.Resolve(ScopedName(trb))
			return HandleDeclType[encoderFunc](
				ast.Decl.Type.Branch,
				func(dtb DeclTypeBranch_Struct_) (encoderFunc, error) {
					return structEncodeBinding(dres, Struct(dtb), texpr.Parameters, boundTypeParams), nil
				},
				func(dtb DeclTypeBranch_Union_) (encoderFunc, error) {
					// union := Union(dtb)
					return nil, nil
				},
				func(dtb DeclTypeBranch_Type_) (encoderFunc, error) {
					return typedefEncodeBinding(dres, TypeDef(dtb), texpr.Parameters, boundTypeParams), nil
				},
				func(dtb DeclTypeBranch_Newtype_) (encoderFunc, error) {
					return newtypeEncodeBinding(dres, NewType(dtb), texpr.Parameters, boundTypeParams), nil
				},
			)
		},
	))
}

func structEncodeBinding(
	dres Resolver,
	struct_ Struct,
	params []TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	newBoundTypeParams := createBoundTypeParams(dres, struct_.TypeParams, params, boundTypeParams)
	fieldJB := make([]encoderFunc, 0)
	for _, field := range struct_.Fields {
		jb := buildEncodeBinding(dres, field.TypeExpr, newBoundTypeParams)
		fieldJB = append(fieldJB, jb)
	}
	fn := func(e *encodeState, v reflect.Value) {
		next := byte('{')
		for i := range struct_.Fields {
			f := struct_.Fields[i]
			fe := fieldJB[i]
			e.WriteByte(next)
			next = ','
			e.WriteString(`"` + f.SerializedName + `":`)
			fe(e, v.Field(i))
		}
		if next == '{' {
			e.WriteString("{}")
		} else {
			e.WriteByte('}')
		}
	}
	return fn
}

func enumEncodeBinding(
	dres Resolver,
	union_ Union,
	params []TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	return nil
}

func unionEncodeBinding(
	dres Resolver,
	union_ Union,
	params []TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	return nil
}

func newtypeEncodeBinding(
	dres Resolver,
	newtype NewType,
	params []TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	return nil
}

func typedefEncodeBinding(
	dres Resolver,
	typedef TypeDef,
	params []TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	return nil
}

func primitiveEncodeBinding(
	dres Resolver,
	ptype string,
	params []TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	switch ptype {
	case "Int8", "Int16", "Int32", "Int64":
		return func(e *encodeState, v reflect.Value) {
			b := e.AvailableBuffer()
			b = strconv.AppendInt(b, v.Int(), 10)
			e.Write(b)
		}
	case "Word8", "Word16", "Word32", "Word64":
		return func(e *encodeState, v reflect.Value) {
			b := e.AvailableBuffer()
			b = strconv.AppendUint(b, v.Uint(), 10)
			e.Write(b)
		}
	case "Bool":
		return func(e *encodeState, v reflect.Value) {
			b := e.AvailableBuffer()
			b = strconv.AppendBool(b, v.Bool())
			e.Write(b)
		}
	case "Float":
		return float64Encoder
	case "Double":
		return float64Encoder
	case "String":
		return func(e *encodeState, v reflect.Value) {
			e.Write(appendString(e.AvailableBuffer(), v.String(), false))
		}
	case "ByteVector":
		return encodeByteSlice
	case "Void":
		return func(e *encodeState, v reflect.Value) {
			e.WriteString("null")
		}
	case "Json":
		return func(e *encodeState, v reflect.Value) {
			if v.IsZero() {
				e.WriteString("null")
				return
			}
			b, _ := gojson.Marshal(v.Interface())
			e.Write(b)
		}
	case "Vector":
		elementBinding := buildEncodeBinding(dres, params[0], boundTypeParams)
		return func(e *encodeState, v reflect.Value) {
			e.WriteByte('[')
			n := v.Len()
			for i := 0; i < n; i++ {
				if i > 0 {
					e.WriteByte(',')
				}
				elementBinding(e, v.Index(i))
			}
			e.WriteByte(']')
		}
	case "StringMap":
		elementBinding := buildEncodeBinding(dres, params[0], boundTypeParams)
		// TODO depends on struct generated for StringMap
		return func(e *encodeState, v reflect.Value) {
			switch v.Kind() {
			case reflect.Array, reflect.Slice:
				e.WriteByte('{')
				n := v.Len()
				for i := 0; i < n; i++ {
					if i > 0 {
						e.WriteByte(',')
					}
					e.WriteString(`"` + v.String() + `"`)
					e.WriteByte(':')
					elementBinding(e, v.Index(i).Field(1))
				}
				e.WriteByte('}')
			case reflect.Map:
				en := stringMapEncoder{elemEnc: elementBinding}
				en.encode(e, v)
			}
		}
	case "Nullable":
		elementBinding := buildEncodeBinding(dres, params[0], boundTypeParams)
		return func(e *encodeState, v reflect.Value) {
			if v.IsNil() {
				e.WriteString("null")
				return
			}
			// depends on how Nullable is encoded
			elementBinding(e, v.Elem())
		}
	}
	return nil
}

func createBoundTypeParams(
	dresolver Resolver,
	paramNames []string,
	paramTypes []TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) boundEncodeTypeParams {
	result := make(boundEncodeTypeParams)
	for i, paramName := range paramNames {
		result[paramName] = buildEncodeBinding(dresolver, paramTypes[i], boundTypeParams)
	}
	return result
}

type floatEncoder int // number of bits

func (bits floatEncoder) encode(e *encodeState, v reflect.Value) {
	f := v.Float()
	// if math.IsInf(f, 0) || math.IsNaN(f) {
	// 	e.error(&UnsupportedValueError{v, strconv.FormatFloat(f, 'g', -1, int(bits))})
	// }

	// Convert as if by ES6 number to string conversion.
	// This matches most other JSON generators.
	// See golang.org/issue/6384 and golang.org/issue/14135.
	// Like fmt %g, but the exponent cutoffs are different
	// and exponents themselves are not padded to two digits.
	b := e.AvailableBuffer()
	// b = mayAppendQuote(b, opts.quoted)
	abs := math.Abs(f)
	fmt := byte('f')
	// Note: Must use float32 comparisons for underlying float32 value to get precise cutoffs right.
	if abs != 0 {
		if bits == 64 && (abs < 1e-6 || abs >= 1e21) || bits == 32 && (float32(abs) < 1e-6 || float32(abs) >= 1e21) {
			fmt = 'e'
		}
	}
	b = strconv.AppendFloat(b, f, fmt, -1, int(bits))
	if fmt == 'e' {
		// clean up e-09 to e-9
		n := len(b)
		if n >= 4 && b[n-4] == 'e' && b[n-3] == '-' && b[n-2] == '0' {
			b[n-2] = b[n-1]
			b = b[:n-1]
		}
	}
	// b = mayAppendQuote(b, opts.quoted)
	e.Write(b)
}

var (
	float32Encoder = (floatEncoder(32)).encode
	float64Encoder = (floatEncoder(64)).encode
)

func appendString[Bytes []byte | string](dst []byte, src Bytes, escapeHTML bool) []byte {
	dst = append(dst, '"')
	start := 0
	for i := 0; i < len(src); {
		if b := src[i]; b < utf8.RuneSelf {
			if htmlSafeSet[b] || (!escapeHTML && safeSet[b]) {
				i++
				continue
			}
			dst = append(dst, src[start:i]...)
			switch b {
			case '\\', '"':
				dst = append(dst, '\\', b)
			case '\b':
				dst = append(dst, '\\', 'b')
			case '\f':
				dst = append(dst, '\\', 'f')
			case '\n':
				dst = append(dst, '\\', 'n')
			case '\r':
				dst = append(dst, '\\', 'r')
			case '\t':
				dst = append(dst, '\\', 't')
			default:
				// This encodes bytes < 0x20 except for \b, \f, \n, \r and \t.
				// If escapeHTML is set, it also escapes <, >, and &
				// because they can lead to security holes when
				// user-controlled strings are rendered into JSON
				// and served to some browsers.
				dst = append(dst, '\\', 'u', '0', '0', hex[b>>4], hex[b&0xF])
			}
			i++
			start = i
			continue
		}
		// TODO(https://go.dev/issue/56948): Use generic utf8 functionality.
		// For now, cast only a small portion of byte slices to a string
		// so that it can be stack allocated. This slows down []byte slightly
		// due to the extra copy, but keeps string performance roughly the same.
		n := len(src) - i
		if n > utf8.UTFMax {
			n = utf8.UTFMax
		}
		c, size := utf8.DecodeRuneInString(string(src[i : i+n]))
		if c == utf8.RuneError && size == 1 {
			dst = append(dst, src[start:i]...)
			dst = append(dst, `\ufffd`...)
			i += size
			start = i
			continue
		}
		// U+2028 is LINE SEPARATOR.
		// U+2029 is PARAGRAPH SEPARATOR.
		// They are both technically valid characters in JSON strings,
		// but don't work in JSONP, which has to be evaluated as JavaScript,
		// and can lead to security holes there. It is valid JSON to
		// escape them, so we do so unconditionally.
		// See https://en.wikipedia.org/wiki/JSON#Safety.
		if c == '\u2028' || c == '\u2029' {
			dst = append(dst, src[start:i]...)
			dst = append(dst, '\\', 'u', '2', '0', '2', hex[c&0xF])
			i += size
			start = i
			continue
		}
		i += size
	}
	dst = append(dst, src[start:]...)
	dst = append(dst, '"')
	return dst
}

func encodeByteSlice(e *encodeState, v reflect.Value) {
	if v.IsNil() {
		e.WriteString("null")
		return
	}

	s := v.Bytes()
	b := e.AvailableBuffer()
	b = append(b, '"')
	b = base64.StdEncoding.AppendEncode(b, s)
	b = append(b, '"')
	e.Write(b)
}

const hex = "0123456789abcdef"

type stringMapEncoder struct {
	elemEnc encoderFunc
}

func (me stringMapEncoder) encode(e *encodeState, v reflect.Value) {
	if v.IsNil() {
		e.WriteString("{}")
		return
	}
	e.WriteByte('{')

	// Extract and sort the keys.
	var (
		sv  = make([]reflectWithString, v.Len())
		mi  = v.MapRange()
		err error
	)
	for i := 0; mi.Next(); i++ {
		if sv[i].ks, err = resolveKeyName(mi.Key()); err != nil {
			// e.error(fmt.Errorf("json: encoding error for type %q: %q", v.Type().String(), err.Error()))
			panic(fmt.Errorf("json: encoding error for type %q: %q", v.Type().String(), err.Error()))

		}
		sv[i].v = mi.Value()
	}
	slices.SortFunc(sv, func(i, j reflectWithString) int {
		return strings.Compare(i.ks, j.ks)
	})

	for i, kv := range sv {
		if i > 0 {
			e.WriteByte(',')
		}
		e.Write(appendString(e.AvailableBuffer(), kv.ks, false))
		e.WriteByte(':')
		me.elemEnc(e, kv.v)
	}
	e.WriteByte('}')
}

type reflectWithString struct {
	v  reflect.Value
	ks string
}

func resolveKeyName(k reflect.Value) (string, error) {
	if k.Kind() == reflect.String {
		return k.String(), nil
	}
	if tm, ok := k.Interface().(encoding.TextMarshaler); ok {
		if k.Kind() == reflect.Pointer && k.IsNil() {
			return "", nil
		}
		buf, err := tm.MarshalText()
		return string(buf), err
	}
	switch k.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return strconv.FormatInt(k.Int(), 10), nil
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return strconv.FormatUint(k.Uint(), 10), nil
	}
	panic("unexpected map key type")
}