package goutil

import (
	"encoding/hex"
	"testing"

	"github.com/matryer/is"
)

func TestXORBytes(t *testing.T) {
	is := is.New(t)

	a, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	is.NoErr(err)

	b, err := hex.DecodeString("686974207468652062756c6c277320657965")
	is.NoErr(err)

	out, err := XORBytes(a, b)
	is.NoErr(err)

	expected, err := hex.DecodeString("746865206b696420646f6e277420706c6179")
	is.NoErr(err)

	is.Equal(expected, out)
}
