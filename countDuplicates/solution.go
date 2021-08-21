package countDuplicates

import "strings"

func DuplicateCount(s1 string) int {
	toLower := strings.ToLower(s1)
	charsMap := make(map[rune]int, len(toLower))

	for _, r := range toLower {
		charsMap[r] +=1
	}

	var counter int
	for _, v := range charsMap {
		if v > 1 {
			counter += 1
		}
	}

	return counter
}
