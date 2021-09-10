package integersRecreation

import "math"

func getDivisors(n int) []int {
	res := make([]int, 0)
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			res = append(res, i)
		}
	}

	return res
}

func sum(c []int) float64 {
	res := 0.0
	for _, v := range c {
		res += math.Pow(float64(v), 2)
	}

	return res
}

func ListSquared(m, n int) [][]int {
	res := make([][]int, 0)
	for i := m; i <= n; i++ {
		divisors := getDivisors(i)
		sumVal := sum(divisors)
		square := math.Sqrt(sumVal)
		if math.Trunc(square)*math.Trunc(square) == sumVal {
			res = append(res, []int{i, int(sumVal)})
		}
	}

	return res
}
