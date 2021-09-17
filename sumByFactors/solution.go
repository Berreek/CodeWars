package sumByFactors

import (
	"fmt"
	"math"
	"sort"
	"strings"
)

func max(c []int) int {
	r := float64(0)
	for _, v := range c {
		abs := math.Abs(float64(v))
		if abs > r {
			r = abs
		}

	}

	return int(r)
}

func sieveOfEratosthenes(max int) []int {
	primes := make([]int, 0, max + 1)
	b := make([]bool, max + 1)

	for i := 2; i <= max; i++ {
		if b[i] == true {
			continue
		}
		primes = append(primes, i)
		for k := i * i; k <= max; k += i {
			b[k] = true
		}
	}
	sort.Ints(primes)

	return primes
}

func calculateSums(input, primes []int) map[int]int {
	r := make(map[int]int, len(primes))

	for _, p := range primes {
		sum := 0
		add := false
		for _, v := range input {
			if v%p == 0 {
				sum += v
				add = true
			}
		}
		if add {
			r[p] = sum
		}
	}
	return r
}

func SumOfDivided(lst []int) string {
	m := max(lst)
	primes := sieveOfEratosthenes(m)
	sums := calculateSums(lst, primes)

	sb := strings.Builder{}
	for _, p := range primes {
		if v, ok := sums[p]; ok {
			sb.WriteString(fmt.Sprintf("(%v %v)", p, v))
		}
	}

	return sb.String()
}
