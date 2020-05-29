package goutil

// MinInt returns the minimum value of an int slice, or 0 if the slice is empty.
func MinInt(vals []int) int {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}

	min := vals[0]

	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

// MinInt8 returns the minimum value of an int8 slice, or 0 if the slice is empty.
func MinInt8(vals []int8) int8 {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}

	min := vals[0]

	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

// MinInt16 returns the minimum value of an int16 slice, or 0 if the slice is empty.
func MinInt16(vals []int16) int16 {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}

	min := vals[0]

	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

// MinInt32 returns the minimum value of an int32 slice, or 0 if the slice is empty.
func MinInt32(vals []int32) int32 {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}

	min := vals[0]

	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}

	return min
}

// MinInt64 returns the minimum value of an int64 slice, or 0 if the slice is empty.
func MinInt64(vals []int64) int64 {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}

	min := vals[0]

	for _, v := range vals[1:] {
		if v < min {
			min = v
		}
	}

	return min
}
