package championship

import (
	"errors"

	"github.com/chrisviana/golang-champions/athlete"
	"github.com/chrisviana/golang-champions/game"
)

type Championship struct {
	Name          string
	Start         string
	End           string
	NumberAthlete int
	Athletes      []*athlete.Athlete
	HashCode      string
	Games         []*game.Game
}

func NewChampionship(name, start, end, hashCode string, numberAthlete int) *Championship {
	return &Championship{
		Name:          name,
		Start:         start,
		End:           end,
		NumberAthlete: numberAthlete,
		Athletes:      make([]*athlete.Athlete, 0),
		Games:         make([]*game.Game, 0),
		HashCode:      hashCode,
	}
}

func (c *Championship) AddAthlete(athlete *athlete.Athlete) error {
	if len(c.Athletes) >= c.NumberAthlete {
		return errors.New("número máximo de atletas atingido")
	}

	c.Athletes = append(c.Athletes, athlete)
	return nil
}

func (c *Championship) CreateGame(athlete1Index, athlete2Index int) error {
	if athlete1Index >= len(c.Athletes) || athlete2Index >= len(c.Athletes) {
		return errors.New("indice de atleta inválido")
	}

	game := game.NewGame(c.Athletes[athlete1Index], c.Athletes[athlete2Index])
	c.Games = append(c.Games, game)
	return nil
}
