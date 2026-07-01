package goutil

import "errors"

// XORBytes returns the result of XORing together a and b, byte by byte.
// It returns an error if a and b are not the same length.
func XORBytes(a, b []byte) ([]byte, error) {
	if len(a) != len(b) {
		return nil, errors.New("buffers must be same length")
	}

	res := make([]byte, len(a))
	for i := range a {
		res[i] = a[i] ^ b[i]
	}

	return res, nil
}
