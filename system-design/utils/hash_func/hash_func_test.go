package hash_func

import (
	"testing"
)

// Table driven test, where we maintain different values to be asserted and apply the logic
type args struct {
	data string
}

var tc = []struct {
	name     string
	args     args
	expected string
}{
	{name: "valid hash output", args: args{"Hey folk"}, expected: "6263ffa08283fb6eb2b1922d782bfc94fd951d1d57b24a1bf1437fde2563217a"},
	{name: "empty args", args: args{""}, expected: ""},
}

func TestHashFunction(t *testing.T) {
	for _, testCase := range tc {
		name := testCase.name
		t.Run(name, func(t *testing.T) {
			if got := HashFunction(testCase.args.data); got != testCase.expected {
				t.Errorf("HashFunction() =%v, expected= %v", got, testCase.expected)
			}
		})
	}
}
