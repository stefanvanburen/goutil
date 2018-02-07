package goutil

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXORBytes(t *testing.T) {
	a, err := hex.DecodeString("1c0111001f010100061a024b53535009181c")
	if err != nil {
		t.Error("Unable to decode string")
	}

	b, err := hex.DecodeString("686974207468652062756c6c277320657965")
	if err != nil {
		t.Error("Unable to decode string")
	}

	out, err := XORBytes(a, b)
	if err != nil {
		t.Error("Unable to XOR input strings")
	}

	expected, err := hex.DecodeString("746865206b696420646f6e277420706c6179")
	if err != nil {
		t.Error("Unable to decode string")
	}

	assert.Equal(t, expected, out, "Should match")
}
