package edityaml

import (
	"encoding/json"
	"fmt"
	"math/big"
	"reflect"
	"regexp"
	"strings"

	"github.com/jbrekelmans/go-edityaml/internal"
	"github.com/jbrekelmans/go-edityaml/internal/pathsyntax"
)

var regexpSimplePathItem = regexp.MustCompile(`^[a-zA-Z][a-zA-Z0-9]*$`)

// Path is a sequence of keys that locates a value within a JSON value / YAML document.
// For example, given a JSON object {"a":{"b":1}}, the path .a.b locates the integer 1.
//
// The empty path is valid.
// For example, given a JSON object {"c":2}, the empty path locates the object itself.
type Path []any

// MustParsePath is a wrapper for ParsePath that panics instead of returning an error.
func MustParsePath(path string) Path {
	pathParsed, err := ParsePath(path)
	if err != nil {
		panic(err)
	}
	return pathParsed
}

// ParsePath parses a path.
// Syntax: zero or more repetions of an "index expression" or a "member access expression".
// Example 1: the syntax .a["b"]["c"].d parses into a path that is a 4-sequence of string keys: a, b, c, d.
// Example 2: the syntax .a[0] parses into a path that is a 2-sequence: string a, integer 0.
// The empty string parses into the empty path.
// Valid strings are either empty, or start with a "." or "[" character.
// "member access expressions" must start with an ASCII letter, and can only consist of ASCII alphanumeric characters.
// Example 3: the syntax .this😀Not😀ASCII is invalid.
// Example 4: the syntax .0DoesNotStartWithLetter is invalid.
// Integer keys must use the "index expression" syntax. Leading zeros are not allowed.
// Example 5: the syntax [1][2] parses into a path that is a 2-sequence: integer 1, integer 2.
// Example 6: the syntax [00] is invalid (leading zeros).
// Example 7: the syntax [01] is invalid (leading zeros).
// Example 8: the syntax [+1] is invalid (+ sign not allowed).
// Example 9: the syntax [-01] is invalid.
// Example 10: the syntax [-1] parses into a path that is a 1-sequence: integer -1.
//
// ParsePath only has syntax for string and integer keys. This means there is no syntax for
// the following keys (which are technically allowed in YAML):
// - the null key
// - boolean keys
// - float keys
// - sequence/list keys
// - map keys
// Paths with null/boolean/float keys can be constructed "manually" using Go expression (i.e. Path{nil, true, false, 1.0}) and
// functions in this library will lookup keys of those types.
// _Keys_ that are themselves maps/sequences are not supported at all by this library. The Go expression Path{ []any{1} } cannot be
// used to lookup a key that itself is a 1-sequence of the integer 1 in a map.
//
// An "index expression" must be a decimal integer or a JSON string surrounded by square brackets.
func ParsePath(path string) (Path, error) {
	if path == "" {
		return Path{}, nil
	}
	p := Path{}
	pathRest := path
	for {
		if pathRest[0] == '.' {
			// Find next . or [
			i := strings.IndexAny(pathRest[1:], "[.") + 1
			if i <= 0 {
				i = len(pathRest)
			}
			pathItem := pathRest[1:i]
			if ok := regexpSimplePathItem.MatchString(pathItem); !ok {
				return nil, fmt.Errorf(`path is invalid: %#v`, path)
			}
			p = append(p, pathItem)
			pathRest = pathRest[i:]
		} else if pathRest[0] == '[' {
			var pathItem any
			var err error
			pathItem, pathRest, err = pathsyntax.ParseIndexExpression(pathRest)
			if err != nil {
				return nil, fmt.Errorf(`error parsing index expression: %w`, err)
			}
			p = append(p, pathItem)
		} else {
			return nil, fmt.Errorf(`path is invalid: %#v`, path)
		}
		if len(pathRest) == 0 {
			break
		}
	}
	return Path(p), nil
}

func (p Path) String() string {
	if len(p) == 0 {
		return "<empty path>"
	}
	const dotNullKey = ".<null>"
	var sb strings.Builder
	for i := 0; i < len(p); i++ {
		switch pathItem := p[i].(type) {
		case string:
			if regexpSimplePathItem.MatchString(pathItem) {
				sb.WriteByte('.')
				sb.WriteString(pathItem)
			} else {
				sb.WriteByte('[')
				x, _ := json.Marshal(pathItem)
				sb.Write(x)
				sb.WriteByte(']')
			}
		case int8, int16, int32, int64, int, uint8, uint16, uint32, uint64, uint:
			_, _ = fmt.Fprintf(&sb, "[%d]", pathItem)
		case big.Int:
			_, _ = fmt.Fprintf(&sb, "[%s]", pathItem.String())
		case *big.Int:
			if pathItem == nil {
				sb.WriteString(dotNullKey)
			} else {
				_, _ = fmt.Fprintf(&sb, "[%d]", pathItem)
			}
		case nil:
			sb.WriteString(dotNullKey)
		default:
			if isNil, _ := internal.IsNilSafe(reflect.ValueOf(pathItem)); isNil {
				sb.WriteString(dotNullKey)
			} else {
				// TODO format other types of integers using "index expression" syntax.
				_, _ = fmt.Fprintf(&sb, ".<key of type %T>", p[i])
			}
		}
	}
	return sb.String()
}
