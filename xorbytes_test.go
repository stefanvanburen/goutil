package goutil

import (
	"encoding/hex"
	"slices"
	"testing"
)

func TestXORBytes(t *testing.T) {
	t.Parallel()
	var (
		v1 = "1c0111001f010100061a024b53535009181c"
		v2 = "686974207468652062756c6c277320657965"
		v3 = "746865206b696420646f6e277420706c6179"
	)
	a, err := hex.DecodeString(v1)
	if err != nil {
		t.Errorf("hex.DecodeString(%v) got unexpected err %v", v1, err)
	}
	b, err := hex.DecodeString(v2)
	if err != nil {
		t.Errorf("hex.DecodeString(%v) got unexpected err %v", v2, err)
	}
	got, err := XORBytes(a, b)
	if err != nil {
		t.Errorf("XORBytes(%v, %v) got unexpected err %v", a, b, err)
	}
	want, err := hex.DecodeString(v3)
	if err != nil {
		t.Errorf("hex.DecodeString(%v) got unexpected err %v", v3, err)
	}
	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
