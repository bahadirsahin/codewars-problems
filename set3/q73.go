package set3

type NBAPlayer struct {
	Team string
	Ppg  float64
}

func Q73(playerOne, playerTwo NBAPlayer) float64 {
	return playerOne.Ppg + playerTwo.Ppg
}
