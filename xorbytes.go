package goutil

import "errors"

// XORBytes returns the result of XORing together two byte slices, byte by byte
// It's required that the input slices are the same length
func XORBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("Buffers must be same length")
	}
	var res = make([]byte, len(a))
	for i := 0; i < len(a); i++ {
		res[i] = a[i] ^ b[i]
	}
	return res, nil
}
