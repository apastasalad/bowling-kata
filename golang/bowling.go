package main

const frames = 10

type Game struct {
}

func (game *Game) score() int {
	score := 0
	counter := 0

	for i := 0; i < frames; i++ {
		score++
		counter++
	}

	return score
}

func (game *Game) roll(pins int) {
}
