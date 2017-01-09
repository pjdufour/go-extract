package extract

import (
	"strconv"
)

import "github.com/pjdufour/go-gypsy/yaml"

func Extract(keyChain string, node yaml.Node, fallback interface{}) interface{} {

	match, err := yaml.Child(node, keyChain)
	if err != nil {
		return fallback
	} else {
		scalar, ok := match.(yaml.Scalar)
		if ok {
			i, err := strconv.Atoi(scalar.String())
			if err == nil {
				return i
			}
			f, err := strconv.ParseFloat(scalar.String(), 64)
			if err == nil {
				return f
			}
			b, err := strconv.ParseBool(scalar.String())
			if err == nil {
				return b
			}
			return scalar.String()
		} else {
			return match
		}
	}
}
