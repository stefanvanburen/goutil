package goutil

type matrix [][]int

type costFunc func(matrix, int, int, int) int

// LevenschteinCost is used for best global alignment.
func LevenschteinCost(m matrix, i, j, cost int) int {
	return min(
		m[i-1][j]+1,
		m[i][j-1]+1,
		m[i-1][j-1]+cost,
	)
}

// SmithWatermanCost is used for best local alignment.
func SmithWatermanCost(m matrix, i, j, cost int) int {
	return min(
		0,
		m[i-1][j]+1,
		m[i][j-1]+1,
		m[i-1][j-1]+cost,
	)
}

// EditDistance1 calculates the edit (or Levenschtein) distance between two
// strings by creating an entire backtraceable matrix.
func EditDistance1(s, t string) int {
	m := editDistanceMatrix(s, t, LevenschteinCost)
	if m == nil {
		return 0
	}
	return m[len(s)][len(t)]
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
			v1[j+1] = min(
				v1[j]+1,
				v0[j+1]+1,
				v0[j]+cost,
			)
		}
		for j := 0; j < len(v0); j++ {
			v0[j] = v1[j]
		}
	}
	return v1[len(t)]
}

// editDistanceMatrix creates a backtraceable matrix.
func editDistanceMatrix(s, t string, costFn costFunc) matrix {
	if checkTrivial(s, t) {
		return nil
	}

	y := len(s) + 1
	x := len(t) + 1

	m := make([][]int, y)

	for z := range m {
		m[z] = make([]int, x)
	}

	for i := 1; i < y; i++ {
		m[i][0] = i
	}
	for j := 1; j < x; j++ {
		m[0][j] = j
	}

	var cost int

	for i := 1; i < y; i++ {
		for j := 1; j < x; j++ {
			cost = 0
			if s[i-1] != t[j-1] {
				cost = 1
			}
			m[i][j] = costFn(m, i, j, cost)
		}
	}
	return m
}

func checkTrivial(s, t string) bool {
	if len(s) == 0 || len(t) == 0 {
		return true
	}
	if s == t {
		return true
	}
	return false
}
