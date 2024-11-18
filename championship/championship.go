package championship

import (
	"github.com/chrisviana/golang-champions/athlete"
	"github.com/chrisviana/golang-champions/game"
)

type Championship struct {
	Name          string
	Start         string
	End           string
	NumberAthlete int
	Athlete       athlete.Athelete
	HashCode      string
	Games         []game.Game
}

func NewChampionship(name, start, end, hashCode string, numberAthlete int) *Championship {
	return &Championship{
		Name:          name,
		Start:         start,
		End:           end,
		NumberAthlete: numberAthlete,
		HashCode:      hashCode,
	}
}
