package main

import "testing"

// A Gutter Game will always return a score of 0
func TestGutterGame(t *testing.T) {
	game := Game{}
	game.roll(1)

	if game.score() != 0 {
		t.Errorf("The Gutter Game returned a non-zero score, got: %d, want: %d.", game.score(), 0)
	}
}
