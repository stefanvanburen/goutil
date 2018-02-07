package goutil

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHexToBase64(t *testing.T) {
	HexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	Base64String := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	HexBytes := []byte(HexString)
	Base64Bytes := []byte(Base64String)

	out, _ := HexStringToBase64String(HexString)

	assert.Equal(t, out, Base64String, "Should match")

	out2, _ := HexToBase64(HexBytes)

	assert.Equal(t, out2, Base64Bytes, "Should match")
}
