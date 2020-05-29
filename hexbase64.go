package goutil

import (
	"encoding/base64"
	"encoding/hex"
)

// HexToBase64 converts a slice of hexadecimal bytes to the equivalent base64 slice of bytes.
func HexToBase64(hexbytes []byte) ([]byte, error) {
	bytes := make([]byte, hex.DecodedLen(len(hexbytes)))

	_, err := hex.Decode(bytes, hexbytes)
	if err != nil {
		return nil, err
	}

	base64bytes := make([]byte, base64.StdEncoding.EncodedLen(len(bytes)))
	base64.StdEncoding.Encode(base64bytes, bytes)

	return base64bytes, nil
}

// HexStringToBase64String converts a hexadecimal-encoded string to a base64-encoded string.
func HexStringToBase64String(hexstring string) (string, error) {
	decodedhex, err := hex.DecodeString(hexstring)
	if err != nil {
		return "", err
	}

	return base64.StdEncoding.EncodeToString(decodedhex), nil
}
