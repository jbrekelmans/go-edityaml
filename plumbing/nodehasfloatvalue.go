package plumbing

import (
	"math"

	goyaml "gopkg.in/yaml.v3"
)

func nodeHasFloatValue(node *goyaml.Node, value any) (bool, error) {
	f1, err := getNodeValueAsFloat(node)
	if err != nil {
		return false, err
	}
	var f2 float64
	switch value := value.(type) {
	case float32:
		f2 = float64(value)
	case float64:
		f2 = value
	default:
		return false, nil
	}
	if math.IsNaN(f1) {
		// YAML defines nans to be equal.
		return math.IsNaN(f2), nil
	}
	return f1 == f2, nil
}
