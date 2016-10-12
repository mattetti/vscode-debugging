package greetings_test

import (
	"testing"

	"github.com/mattetti/vscode-debugging/greetings"
)

func TestSayHi(t *testing.T) {
	testCases := []struct {
		input  string
		output string
	}{
		{"matt", "Hi matt"},
	}

	for i, tc := range testCases {
		t.Logf("test case %d\n", i)
		if o := greetings.SayHello(tc.input); o != tc.output {
			t.Fatalf("expected %s\ngot\n%s\n", tc.output, o)
		}
	}
}
