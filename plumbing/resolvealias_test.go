package plumbing

import (
	"testing"

	"github.com/stretchr/testify/assert"
	goyaml "gopkg.in/yaml.v3"
)

func Test_ResolveAlias(t *testing.T) {
	t.Run("Case1", func(t *testing.T) {
		x := &goyaml.Node{
			Kind: goyaml.AliasNode,
		}
		x.Alias = x
		_, err := ResolveAlias(x)
		assert.ErrorContains(t, err, "aborting after resolving")
	})
	t.Run("Case2", func(t *testing.T) {
		x := &goyaml.Node{
			Kind: goyaml.AliasNode,
		}
		_, err := ResolveAlias(x)
		assert.ErrorContains(t, err, "invalid node")
	})
	t.Run("Case3", func(t *testing.T) {
		expected := &goyaml.Node{}
		x := &goyaml.Node{
			Kind:  goyaml.AliasNode,
			Alias: expected,
		}
		actual, err := ResolveAlias(x)
		if assert.NoError(t, err) {
			assert.Same(t, expected, actual)
		}
	})
}
