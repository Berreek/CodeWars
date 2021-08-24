package validParentheses

func ValidParentheses(parens string) bool {
	parensCounter := 0
	for _, c := range parens {
		switch c {
		case '(':
			parensCounter++
		case ')':
			parensCounter--
		}
		if parensCounter < 0 {
			return false
		}
	}

	return parensCounter == 0
}
