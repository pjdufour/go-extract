package extract_test

import (
	"reflect"
	"testing"
)

import (
	"github.com/pjdufour/go-gypsy/yaml"
	"github.com/pjdufour/go-extract/extract"
)

var filename = "test.yml"

var tests = []struct {
	keyChain      string
	fallback      interface{}
	expectedType  string
	expectedValue interface{}
}{
	{"view.zoom", 3, "int", 11},
	{"controls.attribution", false, "bool", true},
	{"controls.legend", false, "bool", false},
	{"foo", "bar", "string", "bar"},
	{"foo", 14, "int", 14},
	{"baselayers.0.title", nil, "string", "OpenStreetMap"},
	{"baselayers[0].title", nil, "string", "OpenStreetMap"},
}

func TestExtract(t *testing.T) {
	f, err := yaml.ReadFile(filename)

	if err != nil {
		t.Error("Could not open test input file ", err)
	}

	for _, x := range tests {

		value := extract.Extract(x.keyChain, f.Root, x.fallback)

		if reflect.TypeOf(value).String() != x.expectedType {
			t.Error("For key chain", x.keyChain, "expected type of", x.expectedType, "got", reflect.TypeOf(value))
		}

		if value != x.expectedValue {
			t.Error("For key chain", x.keyChain, "expected value", x.expectedValue, "got", value)
		}
	}
}
