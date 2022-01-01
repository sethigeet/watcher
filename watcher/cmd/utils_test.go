package cmd_test

import (
	"reflect"
	"testing"

	"github.com/sethigeet/watcher/watcher/cmd"
)

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
