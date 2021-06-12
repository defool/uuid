package uuid

var (
	codeLen int64 = 62
	codeStr       = []byte("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz")
	codeIdx       = map[byte]int64{}

	codeLen32 int64 = 32
	codeStr32       = []byte("0123456789abcdefghijklmnopqrstuv")
	codeIdx32       = map[byte]int64{}
)

func init() {
	for i, v := range codeStr {
		codeIdx[v] = int64(i)
	}
	for i, v := range codeStr32 {
		codeIdx32[v] = int64(i)
	}
}

// Base62Decode decodes int64 to bytes
func Base62Encode(v int64, ret []byte) {
	size := len(ret)
	var i int
	for ; v > 0 && i < size; i++ {
		ret[size-i-1] = codeStr[v%codeLen]
		v /= codeLen
	}
	for ; i < size; i++ {
		ret[size-i-1] = codeStr[0]
	}
}

// Base62Decode decodes bytes to int64
func Base62Decode(bs []byte) (ret int64) {
	var a int64 = 1
	var size = len(bs)
	for i := range bs {
		ret += codeIdx[bs[size-i-1]] * a
		a *= codeLen
	}
	return
}

// Base32Encode decodes int64 to bytes
func Base32Encode(v int64, ret []byte) {
	size := len(ret)
	var i int
	for ; v > 0 && i < size; i++ {
		ret[size-i-1] = codeStr32[v%codeLen32]
		v /= codeLen32
	}
	for ; i < size; i++ {
		ret[size-i-1] = codeStr32[0]
	}
}

// Base32Decode decodes bytes to int64
func Base32Decode(bs []byte) (ret int64) {
	var a int64 = 1
	var size = len(bs)
	for i := range bs {
		ret += codeIdx32[bs[size-i-1]] * a
		a *= codeLen32
	}
	return
}
