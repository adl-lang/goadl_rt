package goadl

var (
	PrimitiveMap = map[string]string{
		"Int8":       "int8",
		"Int16":      "int16",
		"Int32":      "int32",
		"Int64":      "int64",
		"Word8":      "uint8",
		"Word16":     "uint16",
		"Word32":     "uint32",
		"Word64":     "uint64",
		"Bool":       "bool",
		"Float":      "float64",
		"Double":     "float64",
		"String":     "string",
		"ByteVector": "[]byte",
		"Void":       "struct{}",
		"Json":       "interface{}",
		// "`Vector<T>`":    0,
		// "`StringMap<T>`": 0,
		// "`Nullable<T>`":  0,
	}
)
