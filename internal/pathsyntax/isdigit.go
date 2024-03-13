package pathsyntax

func isDigit(c byte) bool {
	return uint(c-'0') <= 9
}
