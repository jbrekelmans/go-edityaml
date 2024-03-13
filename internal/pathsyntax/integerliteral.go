package pathsyntax

import (
	"errors"
	"math/big"
)

// parseIntegerLiteral parses a decimal integer.
// Leading zeros are not allowed.
// "+0" and "-0" is not allowed.
// text[0] is assumes to be "-" or an ASCII digit.
func parseIntegerLiteral(text string) (value *big.Int, textRest string, err error) {
	if text[0] == '0' {
		if len(text) > 1 && isDigit(text[1]) {
			err = errors.New("leading zeros")
		} else {
			value = new(big.Int)
		}
		textRest = text[1:]
		return
	}
	i := 0
	if text[0] == '-' {
		if len(text) == 1 || !isDigit(text[1]) {
			err = errors.New(`unexpected character "-"`)
			return
		}
		if text[1] == '0' {
			err = errors.New("leading zeros or signed zero")
			return
		}
		i++
	}
	for {
		i++
		if i == len(text) || !isDigit(text[i]) {
			break
		}
	}
	textRest = text[i:]
	// We already checked the syntax so can ignore the "ok" parameter.
	value, _ = new(big.Int).SetString(text[:i], 10)
	return
}
