package oddInt

func FindOdd(seq []int) (result int) {
	countersMap := make(map[int]int, len(seq))
	for _,r := range seq {
		countersMap[r] +=1
	}

	for k,v := range countersMap {
		if v % 2 != 0 {
			result = k
		}
	}

	return result
}
