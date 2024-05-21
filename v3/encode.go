package goadl

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
	"sync"
	"unicode/utf8"

	adlast "github.com/adl-lang/goadl_rt/v3/sys/adlast"
)

type Encoder[T any] struct {
	w       io.Writer
	binding encoderFunc
}

func NewEncoder[T any](
	w io.Writer,
	texpr ATypeExpr[T],
	dres Resolver,
) *Encoder[T] {
	binding := buildEncodeBinding(dres, texpr.Value, make(boundEncodeTypeParams))
	return &Encoder[T]{
		w:       w,
		binding: binding,
	}
}

func (enc *Encoder[T]) Encode(v T) error {
	es := &encodeState{}
	err := enc.binding(es, reflect.ValueOf(v))
	if err != nil {
		return err
	}
	b := es.Bytes()
	_, err = enc.w.Write(b)
	return err
}

type encodeState struct {
	bytes.Buffer // accumulated output
}

type encoderFunc func(e *encodeState, v reflect.Value) error

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

var encoderCache sync.Map // map[reflect.Type]decodeFunc

func buildEncodeBinding(
	dres Resolver,
	texpr adlast.TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	// taken from golang stdlib src/encoding/json/encode.go
	key := texprKey(texpr)
	if fi, ok := encoderCache.Load(key); ok {
		return fi.(encoderFunc)
	}
	// To deal with recursive types, populate the map with an
	// indirect func before we build it. This type waits on the
	// real func (f) to be ready and then calls it. This indirect
	// func is only used for recursive types.
	var (
		wg sync.WaitGroup
		f  encoderFunc
	)
	wg.Add(1)
	fi, loaded := encoderCache.LoadOrStore(key, encoderFunc(func(e *encodeState, v reflect.Value) error {
		wg.Wait()
		return f(e, v)
	}))
	if loaded {
		return fi.(encoderFunc)
	}

	// Compute the real encoder and replace the indirect func with it.
	f = buildNewEncodeBinding(dres, texpr, boundTypeParams)
	wg.Done()
	encoderCache.Store(key, f)
	return f
}

func buildNewEncodeBinding(
	dres Resolver,
	texpr adlast.TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	return adlast.Handle_TypeRef[encoderFunc](
		texpr.TypeRef.Branch,
		func(primitive string) encoderFunc {
			return primitiveEncodeBinding(dres, primitive, texpr.Parameters, boundTypeParams)
		},
		func(typeParam string) encoderFunc {
			return boundTypeParams[typeParam]
		},
		func(reference adlast.ScopedName) encoderFunc {
			ast := dres.Resolve(reference)
			return adlast.Handle_DeclType[encoderFunc](
				ast.SD.Decl.Type_.Branch,
				func(struct_ adlast.Struct) encoderFunc {
					return structEncodeBinding(dres, struct_, texpr.Parameters, boundTypeParams)
				},
				func(union_ adlast.Union) encoderFunc {
					if isEnum(union_) {
						return enumEncodeBinding(dres, union_, texpr.Parameters, boundTypeParams)
					}
					return unionEncodeBinding(dres, union_, texpr.Parameters, boundTypeParams)
				},
				func(type_ adlast.TypeDef) encoderFunc {
					return typedefEncodeBinding(dres, type_, texpr.Parameters, boundTypeParams)
				},
				func(newtype_ adlast.NewType) encoderFunc {
					return newtypeEncodeBinding(dres, newtype_, texpr.Parameters, boundTypeParams)
				},
				nil,
			)
		},
		nil,
	)
}

func structEncodeBinding(
	dres Resolver,
	struct_ adlast.Struct,
	params []adlast.TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	newBoundTypeParams := createBoundTypeParams(dres, struct_.TypeParams, params, boundTypeParams)
	fieldJB := make([]encoderFunc, 0)
	for _, field := range struct_.Fields {
		jb := buildEncodeBinding(dres, field.TypeExpr, newBoundTypeParams)
		fieldJB = append(fieldJB, jb)
	}
	fn := func(e *encodeState, v reflect.Value) error {
		next := byte('{')
		for i := range struct_.Fields {
			f := struct_.Fields[i]
			fe := fieldJB[i]
			e.WriteByte(next)
			next = ','
			e.WriteString(`"` + f.SerializedName + `":`)
			err := fe(e, v.Field(i))
			if err != nil {
				return err
			}
		}
		if next == '{' {
			e.WriteString("{}")
		} else {
			e.WriteByte('}')
		}
		return nil
	}
	return fn
}

func enumEncodeBinding(
	dres Resolver,
	union_ adlast.Union,
	params []adlast.TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	return func(e *encodeState, v reflect.Value) error {
		// name1 := v.Field(0).Type().Name()
		name1 := v.Type().Name()
		name2 := reflect.TypeOf(v.Field(0).Interface()).Name()
		fmt.Printf("n1:%s n2:%s\n", name1, name2)
		key := name2[len(name1)+1:]
		// key := reflect.TypeOf(v.Field(0).Interface()).Field(0).Tag.Get("branch")
		e.WriteString(`"`)
		e.WriteString(string(key))
		e.WriteString(`"`)
		return nil
	}
}

type boundEncField struct {
	encoderFunc encoderFunc
	field       adlast.Field
}

func unionEncodeBinding(
	dres Resolver,
	union_ adlast.Union,
	params []adlast.TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	encMap := make(map[string]boundEncField)
	for _, f := range union_.Fields {
		encMap[f.SerializedName] = boundEncField{
			buildEncodeBinding(dres, f.TypeExpr, boundTypeParams),
			f,
		}
	}

	return func(e *encodeState, v reflect.Value) error {
		if v.Field(0).IsNil() {
			return fmt.Errorf("cannot encode incomplete value")
		}
		key := reflect.TypeOf(v.Field(0).Interface()).Field(0).Tag.Get("branch")
		e.WriteString(`{"`)
		e.WriteString(string(key))
		e.WriteString(`":`)
		if bf, ok := encMap[key]; ok {
			err := bf.encoderFunc(e, reflect.ValueOf(v.Field(0).Interface()).Field(0))
			if err != nil {
				return err
			}
		} else {
			return fmt.Errorf("missing encoding. key: %v", key)
		}
		e.WriteString(`}`)
		return nil
	}
}

func newtypeEncodeBinding(
	dres Resolver,
	newtype adlast.NewType,
	params []adlast.TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	return buildEncodeBinding(dres, newtype.TypeExpr, boundTypeParams)
}

func typedefEncodeBinding(
	dres Resolver,
	typedef adlast.TypeDef,
	params []adlast.TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	return buildEncodeBinding(dres, typedef.TypeExpr, boundTypeParams)
}

func primitiveEncodeBinding(
	dres Resolver,
	ptype string,
	params []adlast.TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) encoderFunc {
	switch ptype {
	case "Int8", "Int16", "Int32", "Int64":
		return func(e *encodeState, v reflect.Value) error {
			b := e.AvailableBuffer()
			b = strconv.AppendInt(b, v.Int(), 10)
			e.Write(b)
			return nil
		}
	case "Word8", "Word16", "Word32", "Word64":
		return func(e *encodeState, v reflect.Value) error {
			b := e.AvailableBuffer()
			b = strconv.AppendUint(b, v.Uint(), 10)
			e.Write(b)
			return nil
		}
	case "Bool":
		return func(e *encodeState, v reflect.Value) error {
			b := e.AvailableBuffer()
			b = strconv.AppendBool(b, v.Bool())
			e.Write(b)
			return nil
		}
	case "Float":
		return float64Encoder
	case "Double":
		return float64Encoder
	case "String":
		return func(e *encodeState, v reflect.Value) error {
			e.Write(appendString(e.AvailableBuffer(), v.String(), false))
			return nil
		}
	case "ByteVector":
		return encodeByteSlice
	case "Void":
		return func(e *encodeState, v reflect.Value) error {
			e.WriteString("null")
			return nil
		}
	case "Json":
		return func(e *encodeState, v reflect.Value) error {
			if v.IsZero() {
				e.WriteString("null")
				return nil
			}
			b, _ := gojson.Marshal(v.Interface())
			e.Write(b)
			return nil
		}
	case "Vector":
		elementBinding := buildEncodeBinding(dres, params[0], boundTypeParams)
		return func(e *encodeState, v reflect.Value) error {
			e.WriteByte('[')
			n := v.Len()
			for i := 0; i < n; i++ {
				if i > 0 {
					e.WriteByte(',')
				}
				err := elementBinding(e, v.Index(i))
				if err != nil {
					return err
				}
			}
			e.WriteByte(']')
			return nil
		}
	case "StringMap":
		elementBinding := buildEncodeBinding(dres, params[0], boundTypeParams)
		// TODO depends on struct generated for StringMap
		return func(e *encodeState, v reflect.Value) error {
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
					err := elementBinding(e, v.Index(i).Field(1))
					if err != nil {
						return err
					}
				}
				e.WriteByte('}')
			case reflect.Map:
				en := stringMapEncoder{elemEnc: elementBinding}
				return en.encode(e, v)
			}
			return nil
		}
	case "Nullable":
		elementBinding := buildEncodeBinding(dres, params[0], boundTypeParams)
		return func(e *encodeState, v reflect.Value) error {
			if v.IsNil() {
				e.WriteString("null")
				return nil
			}
			// depends on how Nullable is encoded
			return elementBinding(e, v.Elem())
		}
	}
	return nil
}

func createBoundTypeParams(
	dresolver Resolver,
	paramNames []string,
	paramTypes []adlast.TypeExpr,
	boundTypeParams boundEncodeTypeParams,
) boundEncodeTypeParams {
	result := make(boundEncodeTypeParams)
	for i, paramName := range paramNames {
		result[paramName] = buildEncodeBinding(dresolver, paramTypes[i], boundTypeParams)
	}
	return result
}

type floatEncoder int // number of bits

func (bits floatEncoder) encode(e *encodeState, v reflect.Value) error {
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
	return nil
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

func encodeByteSlice(e *encodeState, v reflect.Value) error {
	if v.IsNil() {
		e.WriteString("null")
		return nil
	}

	s := v.Bytes()
	b := e.AvailableBuffer()
	b = append(b, '"')
	b = base64.StdEncoding.AppendEncode(b, s)
	b = append(b, '"')
	e.Write(b)
	return nil
}

const hex = "0123456789abcdef"

type stringMapEncoder struct {
	elemEnc encoderFunc
}

func (me stringMapEncoder) encode(e *encodeState, v reflect.Value) error {
	if v.IsNil() {
		e.WriteString("{}")
		return nil
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
	return nil
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

func isEnum(union adlast.Union) bool {
	for _, field := range union.Fields {
		isv := adlast.Handle_TypeRef[bool](
			field.TypeExpr.TypeRef.Branch,
			func(primitive string) bool {
				return primitive == "Void"
			},
			func(typeParam string) bool {
				return false
			},
			func(reference adlast.ScopedName) bool {
				return false
			},
			nil,
		)
		if !isv {
			return false
		}
	}
	return true
}
