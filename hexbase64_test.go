package goutil

import (
	"slices"
	"testing"
)

func TestHexToBase64(t *testing.T) {
	t.Parallel()
	hexString := "49276d206b696c6c696e6720796f757220627261696e206c696b65206120706f69736f6e6f7573206d757368726f6f6d"
	base64String := "SSdtIGtpbGxpbmcgeW91ciBicmFpbiBsaWtlIGEgcG9pc29ub3VzIG11c2hyb29t"

	gotBase64, err := HexStringToBase64String(hexString)
	if err != nil {
		t.Errorf("HexStringToBase64String(%v) got unexpected err %v", hexString, err)
	}
	if gotBase64 != base64String {
		t.Errorf("HexStringToBase64String(%v) = %v, want %v", hexString, gotBase64, base64String)
	}
	gotBase64Bytes, err := HexToBase64([]byte(hexString))
	if err != nil {
		t.Errorf("HexStringToBase64(%v) got unexpected err %v", hexString, err)
	}
	if want := []byte(base64String); !slices.Equal(gotBase64Bytes, want) {
		t.Errorf("HexStringToBase64(%v) = %v, want %v", []byte(hexString), gotBase64Bytes, want)
	}
}
