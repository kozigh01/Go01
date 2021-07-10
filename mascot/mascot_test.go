package mascot_test

import (
	"testing"

	"github.com/kozigh01/go01/mascot"
)

func TestMascot(t *testing.T) {
	if mascot.BestMascot() != "Go Gopher" {
		t.Fatal("wrong mascot :(")
	}
}