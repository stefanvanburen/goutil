package goutil

func Sum(vals []int) (sum int) {
	for _, v := range vals {
		sum += v
	}
	return
}
