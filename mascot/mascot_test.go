package mascot_test

import (
	"testing"

	"github.com/kozigh01/go01/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatalf("wrong mascot :( - expected '%s' but got '%s'", "Go Gopher", mascot.BestMascot())
	}
}