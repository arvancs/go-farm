package farm

import (
	"encoding/binary"
)

func rotate32(val uint32, shift uint) uint32 {
	// Avoid shifting by 32: doing so yields an undefined result.
	if shift == 0 {
		return val
	}
	return ((val >> shift) | (val << (32 - shift)))
}

func rotate64(val uint64, shift uint) uint64 {
	// Avoid shifting by 64: doing so yields an undefined result.
	if shift == 0 {
		return val
	}
	return ((val >> shift) | (val << (64 - shift)))
}

func fetch32(s []byte, idx int) uint32 {
	return binary.LittleEndian.Uint32(s[idx:])
}

func fetch64(s []byte, idx int) uint64 {
	return binary.LittleEndian.Uint64(s[idx:])
}