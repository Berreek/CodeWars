package smallestSum

import (
	"math"
)

type maxes struct {
	firstMaxIndex  int
	firstMaxValue  int
	secondMaxIndex int
	secondMaxValue int
}

func getMaxes(c []int) maxes {
	m1i := math.MinInt64
	m1v := math.MinInt64
	m2i := math.MinInt64
	m2v := math.MinInt64

	for i, v := range c {
		if v > m1v {
			m2v = m1v
			m2i = m1i
			m1i = i
			m1v = v
		} else if v > m2v && v != m1v {
			m2v = v
			m2i = i
		}
	}

	return maxes{
		firstMaxIndex:  m1i,
		firstMaxValue:  m1v,
		secondMaxIndex: m2i,
		secondMaxValue: m2v,
	}
}

func sum(c []int) int {
	r := 0

	for _, v := range c {
		r += v
	}

	return r
}

func Solution(ar []int) int {
	finished := false
	m := getMaxes(ar)

	for !finished {
		newV := m.firstMaxValue - m.secondMaxValue

		if newV <= 0 {
			finished = true
			break
		}

		ar[m.firstMaxIndex] = newV
		if newV > m.secondMaxValue {
			m = maxes{
				firstMaxIndex:  m.firstMaxIndex,
				firstMaxValue:  newV,
				secondMaxValue: m.secondMaxValue,
				secondMaxIndex: m.secondMaxIndex,
			}
		} else {
			m = getMaxes(ar)
		}
	}

	return sum(ar)
}
