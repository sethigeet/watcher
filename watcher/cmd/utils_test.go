package cmd_test

import (
	"os"
	"reflect"
	"testing"

	"github.com/sethigeet/watcher/watcher/cmd"
)

func setupSuite(t *testing.T) {
	// setup the test
	path := "./a/b/c.d/e.f"
	err := os.MkdirAll(path, 0755)
	if err != nil {
		t.Fatalf("os.MkdirAll: %s", err)
	}
}

func TestExists(t *testing.T) {
	tests := map[string]struct {
		input string
		want  bool
	}{
		"correct relative": {input: "../cmd", want: true},
		"wrong relative":   {input: "../asdf", want: false},
		"correct absolute": {input: "/home", want: true},
		"wrong absolute":   {input: "/asdf", want: false},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := cmd.Exists(test.input)
			if !reflect.DeepEqual(test.want, got) {
				t.Fatalf("expected: %v, got: %v", test.want, got)
			}
		})
	}
}

func TestGlob(t *testing.T) {
	setupSuite(t)

	type wantType struct {
		len  int
		vals []string
	}
	tests := map[string]struct {
		input string
		want  wantType
	}{
		"zero double stars, one match": {input: "./*/*/*.d", want: wantType{len: 1, vals: []string{"a/b/c.d"}}},
		"one double star, two matches": {input: "./a/**/*.*", want: wantType{len: 2, vals: []string{"a/b/c.d", "a/b/c.d/e.f"}}},
		"two double starts, one match": {input: "./**/b/**/*.f", want: wantType{len: 1, vals: []string{"a/b/c.d/e.f"}}},
		"empty":                        {input: "/asdf", want: wantType{len: 0, vals: []string{}}},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			matches, err := cmd.Glob(test.input)
			if err != nil {
				t.Fatalf("Glob: %s", err)
			}
			if len(matches) != test.want.len {
				t.Fatalf("got %d matches, expected %d", len(matches), test.want.len)
			}
			for i, match := range matches {
				wanted := test.want.vals[i]
				if match != wanted {
					t.Fatalf("matched [%s], expected [%s]", match, wanted)
				}
			}

		})
	}
}
