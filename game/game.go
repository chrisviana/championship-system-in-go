package game

import "github.com/chrisviana/golang-champions/athlete"

type Game struct {
	Athlete1 *athlete.Athlete
	Athlete2 *athlete.Athlete
	Score1   int
	Score2   int
}

func NewGame(athlete1, athlete2 *athlete.Athlete) *Game {
	return &Game{
		Athlete1: athlete1,
		Athlete2: athlete2,
		Score1:   0,
		Score2:   0,
	}
}
