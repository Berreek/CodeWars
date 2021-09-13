package weirdPrimeGenerator

import "math"

func gcd(a, b int64) int64 {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func generateElements(n int64) []int64 {
	res := make([]int64, 0, n)
	res = append(res, 7)

	for i := int64(1); i < n; i++ {
		res = append(res, res[i-1]+gcd(i+1, res[i-1]))
	}

	return res
}

func calculateDifferences(seq []int64) []int64 {
	n := len(seq)
	res := make([]int64, 0, n)
	res = append(res, 1)

	for i := 1; i < n; i++ {
		res = append(res, seq[i]-seq[i-1])
	}

	return res
}

func CountOnes(n int64) int {
	c := 0
	seq := calculateDifferences(generateElements(n))
	for _, v := range seq {
		if v == 1 {
			c++
		}
	}

	return c
}

func MaxPn(n int64) int64 {
	seq := calculateDifferences(generateElements(10000 ^ n))

	primes := make([]int64, 0, len(seq))
	foundPrimes := make(map[int64]bool, 0)

	for _, v := range seq {
		if n == int64(len(primes)) {
			break
		}

		if _, ok := foundPrimes[v]; v != 1 && !ok {
			primes = append(primes, v)
			foundPrimes[v] = true
		}
	}

	max := int64(math.MinInt64)
	for i, v := range primes {
		if v > max && int64(i) < n {
			max = v
		}
	}

	return max
}

func AnOverAverage(n int64) int {
	seq := generateElements(10000 ^ n)
	diffs := calculateDifferences(seq)
	res := make([]int64, 0, len(seq))

	for i, v := range seq {
		if diffs[i] != 1 {
			res = append(res, v/(int64(i+1)))
		}
		if int64(len(res)) == n {
			break
		}
	}

	sum := int64(0)
	for _, v := range res {
		sum += v
	}

	return int(sum / n)
}
