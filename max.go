package goutil

// MaxInt returns the maximum value of an int slice, or 0 if the slice is empty.
func MaxInt(vals []int) int {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}

	max := vals[0]

	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

// MaxInt8 returns the maximum value of an int8 slice, or 0 if the slice is empty.
func MaxInt8(vals []int8) int8 {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}

	max := vals[0]

	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

// MaxInt16 returns the maximum value of an int16 slice, or 0 if the slice is empty.
func MaxInt16(vals []int16) int16 {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}

	max := vals[0]

	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

// MaxInt32 returns the maximum value of an int32 slice, or 0 if the slice is empty.
func MaxInt32(vals []int32) int32 {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}

	max := vals[0]

	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}

	return max
}

// MaxInt64 returns the maximum value of an int64 slice, or 0 if the slice is empty.
func MaxInt64(vals []int64) int64 {
	if len(vals) == 0 {
		return 0
	} else if len(vals) == 1 {
		return vals[0]
	}

	max := vals[0]

	for _, v := range vals[1:] {
		if v > max {
			max = v
		}
	}

	return max
}
