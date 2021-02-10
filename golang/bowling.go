package main

const frames = 10

type BowlingGame struct {

	// a game consists of 21 potential rolls.
	// 9 ordinary frames; 2 rolls per frame
	// a final frame where no bonus points are allowed,
	// and a bonus roll if there is a strike or a spare
	rolls [21]int

	// keep track of the current roll
	currentRolls int
}

// What's the score?
func (game *BowlingGame) Score() int {
	score := 0
	frameIndex := 0

	for i := 0; i < frames; i++ {
		if DidWeRollAFourBagger(game, frameIndex) {
			score += AwardAFourBaggerBonus()
			frameIndex += 2
		} else if DidWeRollATurkey(game, frameIndex) {
			score += AwardATurkeyBonus()
			frameIndex += 2
		} else if DidWeRollADouble(game, frameIndex) {
			score += AwardADoubleBonus(game, frameIndex)
			frameIndex += 2
		} else if DidWeRollAStrike(game, frameIndex) {
			score += AwardAStrikeBonus(game, frameIndex)
			frameIndex += 2
		} else if DidWeRollASpare(game, frameIndex) {
			score += AwardASpareBonus(game, frameIndex)
			frameIndex += 2
		} else {
			score += game.rolls[frameIndex] + game.rolls[frameIndex+1]
			frameIndex += 2
		}
	}

	// count the last roll if it's awarded
	score += game.rolls[20]

	return score
}

// Look at the frame and determine if a spare was rolled
func DidWeRollASpare(game *BowlingGame, frameIndex int) bool {
	status := false

	if frameIndex < 19 {
		status = game.rolls[frameIndex]+game.rolls[frameIndex+1] == 10
	}

	return status
}

// Look at the frame and determine if a strike was rolled
func DidWeRollAStrike(game *BowlingGame, frameIndex int) bool {
	return game.rolls[frameIndex] == 10
}

// Look at the frames and determine if a double was rolled
func DidWeRollADouble(game *BowlingGame, frameIndex int) bool {
	status := false
	if frameIndex < 18 {
		status = game.rolls[frameIndex] == 10 && game.rolls[frameIndex+2] == 10
	}
	return status
}

// Look at the frames and determine if a Turkey was rolled
func DidWeRollATurkey(game *BowlingGame, frameIndex int) bool {
	status := false
	if frameIndex < 16 {
		status = game.rolls[frameIndex] == 10 && game.rolls[frameIndex+2] == 10 && game.rolls[frameIndex+4] == 10
	}
	return status
}

// Look at the frames and determine if a Four Bagger was rolled
func DidWeRollAFourBagger(game *BowlingGame, frameIndex int) bool {

	status := false
	if frameIndex < 14 {
		status = game.rolls[frameIndex] == 10 && game.rolls[frameIndex+2] == 10 && game.rolls[frameIndex+3] == 10
	}
	return status
}

// 10 points + the number of pins you knock down for your first attempt at the next frame.
func AwardASpareBonus(game *BowlingGame, frameIndex int) int {
	// a spare always counts as ten points
	spareBonus := 10

	// only award bonus points if the spare occurs in the first 9 frames.
	if frameIndex < 18 {
		spareBonus += game.rolls[frameIndex+2]
	}
	return spareBonus
}

// 10 points + the number of pins you knock down in the entire next frame.
func AwardAStrikeBonus(game *BowlingGame, frameIndex int) int {
	strikeBonus := 10

	// If the strike happens in the first **9** frames, count the next two rolls
	// otherwise, count the bonus roll only
	if frameIndex < 18 {
		strikeBonus += game.rolls[frameIndex+2] + game.rolls[frameIndex+3]
	} else if frameIndex >= 18 {
		strikeBonus += game.rolls[frameIndex+1] + game.rolls[frameIndex+2]
	}

	return strikeBonus
}

// 20 points + the number of pins you knock down in the third frame.
func AwardADoubleBonus(game *BowlingGame, frameIndex int) int {
	doubleBonus := 20

	// If the double happens in the first **8** frames, count the next two rolls
	// otherwise, count the bonus roll only
	if frameIndex < 16 {
		doubleBonus += game.rolls[frameIndex+4] + game.rolls[frameIndex+5]
	} else if frameIndex >= 16 {
		doubleBonus += game.rolls[frameIndex+1] + game.rolls[frameIndex+2]
	}
	return doubleBonus
}

//  30 point bonus
func AwardATurkeyBonus() int {
	turkeyBonus := 30
	return turkeyBonus
}

//  30 point bonus
func AwardAFourBaggerBonus() int {
	fourBaggerBonus := 30
	return fourBaggerBonus
}

// Knock down some pins
func (game *BowlingGame) Roll(pins int) {
	// Record the pins knocked down for this roll
	// if we get a strike, record a 0 and move to the next frame

	// record the first 9 frames differently to the final frame
	if game.currentRolls < 19 {
		game.rolls[game.currentRolls] = pins
		if pins < 10 {
			game.currentRolls++
		} else {
			game.currentRolls += 2
		}
	} else {
		game.rolls[game.currentRolls] = pins
		game.currentRolls++
	}
}
