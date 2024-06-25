package matchmaking

type Player struct {
	Username string
}

func NewPlayer(username string) *Player {
	return &Player{
		Username: username,
	}
}
