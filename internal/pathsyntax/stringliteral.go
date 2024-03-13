package pathsyntax

import (
	"encoding/json"
	"errors"
	"fmt"
	"unicode/utf8"
)

// parseStringLiteral parses a JSON string literal.
// text[0] is assumed to be a double quote.
func parseStringLiteral(text string) (value, textRest string, err error) {
	i := 1
	hasEscapeSequences := false
	for {
		if i == len(text) {
			// Missing double quote.
			err = errors.New(`unterminated string literal`)
			return
		}
		if text[i] == '\\' {
			i++
			if i == len(text) {
				// End-of-string while processing escape sequence.
				err = errors.New(`unterminated string literal`)
				return
			}
			hasEscapeSequences = true
		} else if text[i] == '"' {
			break
		}
		i++
	}
	i++
	textRest = text[i:]
	if !hasEscapeSequences {
		value = text[1 : i-1]
		if utf8.ValidString(value) {
			return
		}
	}
	err = json.Unmarshal([]byte(text[:i]), &value)
	if err != nil {
		err = fmt.Errorf(`string literal is invalid: %w`, err)
	}
	return
}
