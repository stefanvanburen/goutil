package goutil

import "errors"

var ErrBuffersNotSameLength = errors.New("buffers must be same length")

// XORBytes returns the result of XORing together a and b, byte by byte.
// It's required that the input slices are the same length.
func XORBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, ErrBuffersNotSameLength
	}

	res := make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] ^ b[i]
	}

	return res, nil
}
