package main

const frames = 10

type BowlingGame struct {

	// a game consists of 21 potential rolls.
	// 10 frames; 2 rolls per frame and a final roll if there is a strike
	rolls [21]int

	// keep track of the current roll
	currentRolls int
}

// What's the score?
func (game *BowlingGame) Score() int {
	score := 0
	counter := 0

	for i := 0; i < frames; i++ {
		score += game.rolls[counter] + game.rolls[counter+1]

		// fmt.Printf("the score is %v\n", score)
		counter += 2
	}

	return score
}

// Knock down some pins
func (game *BowlingGame) Roll(pins int) {

	// Record the pins knocked down for this roll
	game.rolls[game.currentRolls] = pins
	game.currentRolls++
}
