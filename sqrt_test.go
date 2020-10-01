package goutil

import (
	"testing"

	"github.com/matryer/is"
)

func TestSqrtInt(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		in  int
		out int
	}{
		{4, 2},
		{9, 3},
		{100, 10},
		{144, 12},
	}

	for _, x := range tests {
		is.Equal(SqrtInt(x.in), x.out)
	}
}

func TestSqrtInt64(t *testing.T) {
	is := is.New(t)

	tests := []struct {
		in  int64
		out int64
	}{
		{4, 2},
		{9, 3},
		{100, 10},
		{144, 12},
	}

	for _, x := range tests {
		is.Equal(SqrtInt64(x.in), x.out)
	}
}
