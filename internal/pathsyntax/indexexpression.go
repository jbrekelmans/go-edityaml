package pathsyntax

import (
	"errors"
	"fmt"

	"math/big"
)

// ParseIndexExpression parses an expression like the following:
// ["name-of-property\\with\"special characters"]
// [0]
// The double quoted string must be a JSON string literal.
// path[0] is assumed to be the left square bracket starting the index expression.
func ParseIndexExpression(path string) (pathItem any, pathRest string, err error) {
	if len(path) < 2 {
		err = errors.New(`unterminated index expression`)
		return
	}
	switch {
	case path[1] == '"':
		var value string
		value, pathRest, err = parseStringLiteral(path[1:])
		if err != nil {
			return
		}
		pathItem = value
	case uint(path[1]-'0') <= 9 || path[1] == '-':
		var value *big.Int
		value, pathRest, err = parseIntegerLiteral(path[1:])
		if err != nil {
			err = fmt.Errorf(`invalid integer literal: %w`, err)
			return
		}
		pathItem = value
	default:
		err = errors.New(`"[" character that start index expressions must be immediately followed by a string/integer literal`)
		return
	}
	if pathRest == "" {
		err = errors.New(`unterminated index expression`)
		return
	}
	if pathRest[0] != ']' {
		err = errors.New(`string/integer literal in an index expression must be immediately followed by "]" character`)
		return
	}
	pathRest = pathRest[1:]
	return
}
