package main

import "strings"

// EditDistance1 calculates the edit (or Levenschtein) distance between two
// strings by creating an entire backtrackable matrix
func EditDistance1(s, t string) int {
	if checkTrivial(s, t) {
		return 0
	}

	y := len(s) + 1
	x := len(t) + 1

	matrix := make([][]int, y)

	for z := range matrix {
		matrix[z] = make([]int, x)
	}

	for i := 1; i < y; i++ {
		matrix[i][0] = i
	}
	for j := 1; j < x; j++ {
		matrix[0][j] = j
	}

	for i := 1; i < y; i++ {
		for j := 1; j < x; j++ {
			var cost int
			if s[i-1] != t[j-1] {
				cost = 1
			}
			b := []int{
				matrix[i-1][j] + 1,
				matrix[i][j-1] + 1,
				matrix[i-1][j-1] + cost,
			}
			matrix[i][j] = Min(b)
		}
	}
	return matrix[len(s)][len(t)]
}

// EditDistance2 calculates the edit (or Levenschtein) distance between two
// strings by just using the previous row and the previous values in the
// current row. This saves memory, but requires copying rows forward.
func EditDistance2(s, t string) int {
	if checkTrivial(s, t) {
		return 0
	}

	v0 := make([]int, len(t)+1)
	v1 := make([]int, len(t)+1)

	for i := 0; i < len(v0); i++ {
		v0[i] = i
	}

	for i := 0; i < len(s); i++ {
		v1[0] = i + 1
		for j := 0; j < len(t); j++ {
			var cost int
			if s[i] != t[j] {
				cost = 1
			}
			b := []int{
				v1[j] + 1,
				v0[j+1] + 1,
				v0[j] + cost,
			}
			v1[j+1] = Min(b)
		}
		for j := 0; j < len(v0); j++ {
			v0[j] = v1[j]
		}
	}
	return v1[len(t)]
}

func checkTrivial(s, t string) bool {
	if strings.Compare(s, t) == 0 {
		return true
	}
	if len(s) == 0 || len(t) == 0 {
		return true
	}
	return false
}
