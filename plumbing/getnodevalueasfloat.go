package plumbing

import (
	"fmt"
	"math"
	"strconv"

	goyaml "gopkg.in/yaml.v3"
)

func getNodeValueAsFloat(node *goyaml.Node) (float64, error) {
	if s := node.Value; s != "" {
		switch s {
		case ".nan", ".NaN", ".NAN":
			return math.NaN(), nil
		}
		sign := 1
		if s[0] == '-' {
			sign = -1
			s = s[1:]
		} else if s[0] == '+' {
			s = s[1:]
		}
		switch s {
		case ".inf", ".Inf", ".INF":
			return math.Inf(sign), nil
		}
	}
	f, err := strconv.ParseFloat(node.Value, 64)
	if err != nil {
		return 0.0, fmt.Errorf(`invalid node (shortTag = "!!float", value = %#v)`, node.Value)
	}
	return f, nil
}
