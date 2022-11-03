package repl_test

import (
	repl "chance"
	"testing"
)

func TestRepl(t *testing.T) {
	x := 0
	for {
		x += 1
		repl.Repl()
		if x == 10 {
			break
		}
	}
}