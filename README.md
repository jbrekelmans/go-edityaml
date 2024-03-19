# go-edityaml

`go-edityaml` is a library for modifying YAML represented as an Abstract Syntax Tree (AST).

Useful -:
1. to update YAML from automation _while preserving comments and formatting/style_.
2. for (command line) tooling that updates YAML that is also read/modified directly by humans.

## Getting Started

```bash
go get github.com/jbrekelmans/go-edityaml
```

## About go-edityaml

The AST type is *Node of the library "gopkg.in/yaml.v3". This AST type cannot 100% accurately
represent all YAML, but is very close and will preserve comments.

The YAML specification (https:yaml.org/spec/1.2.2/#canonical-form) is very complex and maps
can have different types of keys. For example:

    ```yaml
	---
	mapExample:
	   key1: value1 # String key, string value

	   ? key2 # String key (using "complex mapping key" syntax)
	   : value2 # String value

	   ? # Null key
	   : 1 # Integer value.

	   ? [1, 2] # Sequence key
	   : 1.0  #float value.

	   ?  key1: v1 # Key that itself is a map.
	      key2: v2
       :  # Null value
    ```

For this reason, this package uses type any / interface{} when looking up keys in maps,
and Golang bool/integer/float/string types intuitively map to the corresponding YAML scalars.
Furthermore, this package does not support looking up keys that are themselves a sequence/map at all
(i.e. Golang slices / maps are not mapped to YAML sequences / maps).
