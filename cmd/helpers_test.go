package cmd

import (
	"testing"
)

func TestHelpers(t *testing.T) {
	op := SanitiseStringInput("hello tom")
	if op != "hello_tom" {
		t.Fatal("Sanitization failed")
	}
	op = SanitiseStringInput("move fast and break things")
	if op != "move_fast_and_break_things" {
		t.Fatal("Sanitization failed")
	}

}
