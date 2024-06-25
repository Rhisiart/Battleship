package matchmaking

type Match struct {
	Navy  *Player
	Seals *Player
}

func NewMatch(navy *Player, seals *Player) *Match {
	return &Match{
		Navy:  navy,
		Seals: seals,
	}
}
