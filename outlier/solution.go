package outlier

func FindOutlier(integers []int) int {
	isOutlierEven := isOutlierEven(integers)

	for _, v := range integers {
		if v%2 == 0 && isOutlierEven {
			return v
		} else if v%2 != 0 && !isOutlierEven {
			return v
		}
	}

	return 0
}

func isOutlierEven(integers []int) bool {
	var evenCounter int
	var oddCounter int

	firstThreeElements := integers[:3]
	for _, v := range firstThreeElements {
		if v%2 == 0 {
			evenCounter++
		} else {
			oddCounter++
		}
	}
	outlierIsEven := evenCounter < oddCounter
	return outlierIsEven
}
