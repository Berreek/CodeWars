package squares

func decomposer(n, remain int64) []int64 {
	if remain == 0 {
		return []int64{n}
	}

	for i := n - 1; i > 0; i-- {
		if remain-i*i >= 0 {
			r := decomposer(i, remain-i*i)
			if len(r) > 0 {
				r = append(r, n)
				return r
			}
		}
	}

	return []int64{}
}

func Decompose(n int64) []int64 {
	r := decomposer(n, n*n)
	if len(r) > 0 {
		return r[:len(r)-1]
	}

	return r
}
