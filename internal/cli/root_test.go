package cli

import "testing"

func TestRootCommandShape(t *testing.T) {
	cmd := newRootCmd()
	if cmd.Name() != "hst" {
		t.Fatalf("name = %q", cmd.Name())
	}

	sub := map[string]bool{}
	for _, c := range cmd.Commands() {
		sub[c.Name()] = true
	}

	for _, expected := range []string{"remove", "sync", "completion"} {
		if !sub[expected] {
			t.Fatalf("missing subcommand %q", expected)
		}
	}
}
