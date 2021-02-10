package main

import (
	"testing"
)

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

// Spare in first roll, one pin down in each other roll, (score = 11(spare) + 18 + 1 = 30)
func TestSpareInFirstRoll(t *testing.T) {
	game := BowlingGame{}
	LetsRollASpare(&game)
	LetsRoll(&game, 19, 1)

	if game.Score() != 30 {
		t.Errorf("Spare in first roll, one pin down in each other roll, got: %d, want: %d.", game.Score(), 30)
	}
}

// Spare in last roll, one pin down in each other roll, (score = 18 + 10  = 28)
func TestSpareInlastRoll(t *testing.T) {
	game := BowlingGame{}
	LetsRoll(&game, 18, 1)
	LetsRollASpare(&game)

	if game.Score() != 28 {
		t.Errorf("Spare in last roll, one pin down in each other roll, got: %d, want: %d.", game.Score(), 28)
	}
}

//Strike in first roll, one pin down in each other roll, (score = 12 + 19 = 31)
func TestStrikeInFirstRoll(t *testing.T) {
	game := BowlingGame{}
	LetsRollAStrike(&game)
	LetsRoll(&game, 19, 1)

	if game.Score() != 31 {
		t.Errorf("Strike in first roll, one pin down in each other roll, got: %d, want: %d.", game.Score(), 31)
	}
}

//Strike in first roll, one pin down in each other roll, (score = 10 + 1 + 1 + 18 = 30)
//Strike in last roll, one pin down in each other roll, (score = 18 + 10 + 1 + 1 = 30)
//Golden game = all strikes (score = 300)

// Roll the ball and knock down some pins
func LetsRoll(g *BowlingGame, numberOfRolls, pinsKnockedDown int) {
	for i := 0; i < numberOfRolls; i++ {
		g.Roll(pinsKnockedDown)
	}
}

// Roll the ball and get a spare
func LetsRollASpare(g *BowlingGame) {
	g.Roll(8)
	g.Roll(2)
}

// Roll the ball and get a strike
func LetsRollAStrike(g *BowlingGame) {
	g.Roll(10)
}
