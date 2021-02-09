package main

import "testing"

func TestGutterGame(t *testing.T) {
	game := Game{}
	game.roll(1)
	t.Fail()
}
