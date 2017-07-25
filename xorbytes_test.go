package goutil

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestXORBytes(t *testing.T) {
	a := []byte("1c0111001f010100061a024b53535009181c")
	fmt.Println(a)
	b := []byte("686974207468652062756c6c277320657965")
	fmt.Println(b)
	expected := []byte("746865206b696420646f6e277420706c6179")
	fmt.Println(len(expected))

	out, _ := XORBytes(a, b)

	assert.Equal(t, expected, out, "Should match")

}
