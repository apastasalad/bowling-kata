package main

import "testing"

// A Gutter Game will always return a score of 0
func TestGutterGame(t *testing.T) {
	game := BowlingGame{}
	LetsRoll(&game, 20, 0)

	if game.Score() != 0 {
		t.Errorf("The Gutter Game returned a non-zero score, got: %d, want: %d.", game.Score(), 0)
	}
}

// Knock down one pin each roll, (score = 20)
func TestOnePinAtATime(t *testing.T) {
	game := BowlingGame{}
	LetsRoll(&game, 20, 1)

	if game.Score() != 20 {
		t.Errorf("The One Pin at a Time Game returned an incorrect, got: %d, want: %d.", game.Score(), 20)
	}
}

// Spare in first roll, one pin down in each other roll, (score = 10 + 1 + 18 = 29)
func TestSpareInFirstRoll(t *testing.T) {
	game := BowlingGame{}
	LetsRoll(&game, 1, 10)
	LetsRoll(&game, 18, 1)

	if game.Score() != 29 {
		t.Errorf("Spare in first roll, one pin down in each other roll, got: %d, want: %d.", game.Score(), 29)
	}
}

//
//Spare in last roll, one pin down in each other roll, (score = 18 + 10 + 1 = 29)
//Strike in first roll, one pin down in each other roll, (score = 10 + 1 + 1 + 18 = 30)
//Strike in last roll, one pin down in each other roll, (score = 18 + 10 + 1 + 1 = 30)
//Golden game = all strikes (score = 300)

// Roll the ball and knock down some pins
func LetsRoll(g *BowlingGame, numberOfRolls, pinsKnockedDown int) {
	for i := 0; i < numberOfRolls; i++ {
		g.Roll(pinsKnockedDown)
	}
}
