package shquote

import (
	"fmt"
	"path"
	"testing"
)

func TestQuote(t *testing.T) {
	tests := []struct {
		input, output string
	}{
		{"", "''"},
		{"abc", "abc"},
		{"$x", "'$x'"},
		{"\"x\"", "'\"x\"'"},
		{"'x'", "''\"'\"'x'\"'\"''"},
	}

	for _, tst := range tests {
		got := Quote(tst.input)
		if got != tst.output {
			t.Errorf("Quote(%q) produced %q; expected %q", tst.input, got, tst.output)
		}
	}
}

func TestQuoteList(t *testing.T) {
	tests := []struct {
		input  []string
		output string
	}{
		{[]string{}, ""},
		{[]string{"a", "b"}, "a b"},
		{[]string{"$", "b"}, "'$' b"},
		{[]string{"x", "a b c", "y"}, "x 'a b c' y"},
	}

	for _, tst := range tests {
		got := QuoteList(tst.input)
		if got != tst.output {
			t.Errorf("Quote(%q) produced %q; expected %q", tst.input, got, tst.output)
		}
	}
}

func ExampleQuote() {
	projroot := "/opt/projects"
	projname := "my project"
	fmt.Printf("GOPATH=%s\n", Quote(path.Join(projroot, projname)))
	// Output: GOPATH='/opt/projects/my project'
}

func ExampleQuoteList() {
	args := []string{"(", "x=10", "-o", "y=20", ")"}
	fmt.Printf("calc %s\n", QuoteList(args))
	// Output: calc '(' x=10 -o y=20 ')'
}
