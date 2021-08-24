package bookseller

import (
	"fmt"
	"strconv"
	"strings"
)

func StockList(listArt []string, listCat []string) string {
	if len(listArt) == 0 || len(listCat) == 0 {
		return ""
	}

	categoriesQuantity := make(map[string]int, len(listArt))
	for _, book := range listArt {
		quantity, _ := strconv.Atoi(strings.Split(book, " ")[1])
		categoriesQuantity[book[0:1]] += quantity
	}

	elements := make([]string, len(listCat))
	for i, category := range listCat {
		elements[i] = fmt.Sprintf("(%s : %v)", category, categoriesQuantity[category])
	}

	return strings.Join(elements, " - ")
}