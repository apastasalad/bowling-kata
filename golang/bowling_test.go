package main

import "testing"

// A Gutter Game will always return a score of 0
func TestGutterGame(t *testing.T) {
	game := BowlingGame{}
	PlayAGame(&game, 20, 0)

	if game.Score() != 0 {
		t.Errorf("The Gutter Game returned a non-zero score, got: %d, want: %d.", game.Score(), 0)
	}
}

// A Gutter Game will always return a score of 0
func TestOnePinAtATime(t *testing.T) {
	game := BowlingGame{}
	PlayAGame(&game, 20, 1)

	if game.Score() != 20 {
		t.Errorf("The One Pin at a Time Game returned an incorrect, got: %d, want: %d.", game.Score(), 20)
	}
}

func PlayAGame(g *BowlingGame, numberOfRolls, pinsKnockedDown int) {
	for i := 0; i < numberOfRolls; i++ {
		g.Roll(pinsKnockedDown)
	}
}
