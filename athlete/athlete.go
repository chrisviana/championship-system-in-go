package athlete

type Athlete struct {
	Name            string
	PointsEarned    int
	NumberVictories int
	NumberDefeats   int
	StartingBalance int
	GamesPlayed     int
	Performance     int
}

func NewAthlete(name string) *Athlete {
	return &Athlete{
		Name:            name,
		PointsEarned:    0,
		NumberVictories: 0,
		NumberDefeats:   0,
		StartingBalance: 0,
		GamesPlayed:     0,
		Performance:     0,
	}
}
