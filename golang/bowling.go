package main

import "fmt"

const frames = 10

type BowlingGame struct {

	// a game consists of 21 potential rolls.
	// 10 frames; 2 rolls per frame and a final roll if there is a strike or a spare
	rolls [21]int

	// keep track of the current roll
	currentRolls int
}

// What's the score?
func (game *BowlingGame) Score() int {
	fmt.Printf("%v", game.rolls)
	score := 0
	frameCounter := 0

	for i := 0; i < frames; i++ {

		if DidWeRollAStrike(game, frameCounter) {
			score += AwardAStrikeBonus(game, frameCounter)
			fmt.Printf("%v, ", score)
			frameCounter += 2
		} else if DidWeRollASpare(game, frameCounter) {
			score += AwardASpareBonus(game, frameCounter)
			fmt.Printf("%v, ", score)
			frameCounter += 2
		} else {
			score += game.rolls[frameCounter] + game.rolls[frameCounter+1]
			fmt.Printf("%v, ", score)
			frameCounter += 2
		}
	}

	// count the last roll if it's awarded
	score += game.rolls[20]

	fmt.Printf("FINAL SCORE %v, ", score)

	return score
}

// Look at the frame and determine if a spare was rolled
func DidWeRollASpare(game *BowlingGame, frameIndex int) bool {
	return game.rolls[frameIndex]+game.rolls[frameIndex+1] == 10
}

// Look at the frame and determine if a strike was rolled
func DidWeRollAStrike(game *BowlingGame, frameIndex int) bool {
	return game.rolls[frameIndex] == 10
}

// 10 points + the number of pins you knock down in the entire next frame.
func AwardAStrikeBonus(game *BowlingGame, frameIndex int) int {
	strikeBonus := 10

	// If the strike happens in the first 9 frames, count the next two rolls
	// otherwise, count the bonus roll only
	if frameIndex < 18 {
		strikeBonus += game.rolls[frameIndex+2] + game.rolls[frameIndex+3]
	} else if frameIndex >= 18 {
		strikeBonus += game.rolls[frameIndex+1] + game.rolls[frameIndex+2]
	}

	return strikeBonus
}

// 10 points + the number of pins you knock down for your first attempt at the next frame.
func AwardASpareBonus(game *BowlingGame, frameIndex int) int {
	spareBonus := 10
	if frameIndex < 19 {
		spareBonus += game.rolls[frameIndex+2]
	}
	return spareBonus
}

// Knock down some pins
func (game *BowlingGame) Roll(pins int) {
	// Record the pins knocked down for this roll
	// if we get a strike, record a 0 and move to the next frame
	game.rolls[game.currentRolls] = pins
	if pins < 10 {
		game.currentRolls++
	} else {
		game.currentRolls += 2
	}

}
