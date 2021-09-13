package friendCheating

func RemovNb(m uint64) [][2]uint64 {
	res := make([][2]uint64, 0, m)
	sum := m * (m + 1) / 2

	for i := m / 2; i <= m; i++ {
		factor := (sum -i) / (i + 1)
		prod := i * factor
		if prod == sum - i - factor{
			res = append(res, [2]uint64{i, factor})
		}
	}

	if len(res) == 0 {
		return nil
	}

	return res
}
